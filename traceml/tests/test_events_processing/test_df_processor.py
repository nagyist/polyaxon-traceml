import numpy as np
import pandas as pd
import pytest

from pandas.testing import assert_series_equal
from random import shuffle
from unittest import TestCase

from clipped.utils.units import to_percentage

from traceml.processors import df_processors


@pytest.mark.processors_mark
class DataFrameSummaryTest(TestCase):
    def setUp(self):
        self.size = 1000
        missing = [np.nan] * (self.size // 10) + list(range(10)) * (
            (self.size - self.size // 10) // 10
        )
        shuffle(missing)

        self.types = pd.Index(
            [
                df_processors.DFTypes.NUMERIC,
                df_processors.DFTypes.BOOL,
                df_processors.DFTypes.CATEGORICAL,
                df_processors.DFTypes.CONSTANT,
                df_processors.DFTypes.UNIQUE,
                df_processors.DFTypes.DATE,
            ],
            name="types",
        )

        self.columns = [
            "dbool1",
            "dbool2",
            "duniques",
            "dcategoricals",
            "dnumerics1",
            "dnumerics2",
            "dnumerics3",
            "dmissing",
            "dconstant",
            "ddates",
        ]

        self.df = pd.DataFrame(
            dict(
                dbool1=np.random.choice([0, 1], size=self.size),
                dbool2=np.random.choice(["a", "b"], size=self.size),
                duniques=["x{}".format(i) for i in range(self.size)],
                dcategoricals=[
                    "a" if i % 2 == 0 else "b" if i % 3 == 0 else "c"
                    for i in range(self.size)
                ],
                dnumerics1=range(self.size),
                dnumerics2=range(self.size, 2 * self.size),
                dnumerics3=list(range(self.size - self.size // 10))
                + list(range(-self.size // 10, 0)),
                dmissing=missing,
                dconstant=["a"] * self.size,
                ddates=pd.date_range("2010-01-01", periods=self.size, freq="1M"),
            )
        )
        self.column_stats = df_processors.get_df_column_stats(self.df)
        self.columns_types = df_processors.get_df_columns_types(self.column_stats)

    def test_get_columns_works_as_expected(self):
        assert (
            len(df_processors.get_df_columns(self.df, df_processors.DFUsage.ALL)) == 10
        )

        assert (
            len(
                df_processors.get_df_columns(
                    self.df,
                    df_processors.DFUsage.INCLUDE,
                    ["dnumerics1", "dnumerics2", "dnumerics3"],
                )
            )
            == 3
        )

        assert (
            len(
                df_processors.get_df_columns(
                    self.df,
                    df_processors.DFUsage.EXCLUDE,
                    ["dnumerics1", "dnumerics2", "dnumerics3"],
                )
            )
            == 7
        )

    def test_column_types_works_as_expected(self):
        result = self.columns_types[self.types]
        expected = pd.Series(
            index=self.types, data=[4, 2, 1, 1, 1, 1], name=result.name
        )[self.types]
        assert_series_equal(result, expected)

    def test_column_stats_works_as_expected(self):
        self.assertTupleEqual(self.column_stats.shape, (5, 10))

        # counts
        expected = pd.Series(
            index=self.columns, data=self.size, name="counts", dtype="object"
        )
        expected["dmissing"] -= 100
        assert_series_equal(
            self.column_stats[self.columns].loc["counts"], expected[self.columns]
        )

        # uniques
        expected = pd.Series(
            index=self.columns, data=self.size, name="uniques", dtype="object"
        )
        expected[["dbool1", "dbool2"]] = 2
        expected[["dcategoricals"]] = 3
        expected[["dconstant"]] = 1
        expected[["dmissing"]] = 10
        assert_series_equal(
            self.column_stats[self.columns].loc["uniques"].sort_index(),
            expected[self.columns].sort_index(),
            check_dtype=False,
        )

        # missing
        expected = pd.Series(index=self.columns, data=0, name="missing", dtype="object")
        expected[["dmissing"]] = 100
        assert_series_equal(
            self.column_stats[self.columns].loc["missing"],
            expected[self.columns],
            check_dtype=False,
        )

        # missing_perc
        expected = pd.Series(
            index=self.columns, data=["0%"] * 10, name="missing_perc", dtype="object"
        )

        expected[["dmissing"]] = "10%"
        assert_series_equal(
            self.column_stats[self.columns].loc["missing_perc"], expected[self.columns]
        )

        # types
        expected = pd.Series(
            index=self.columns, data=[np.nan] * 10, name="types", dtype="object"
        )

        expected[["dbool1", "dbool2"]] = df_processors.DFTypes.BOOL
        expected[["dcategoricals"]] = df_processors.DFTypes.CATEGORICAL
        expected[["dconstant"]] = df_processors.DFTypes.CONSTANT
        expected[["ddates"]] = df_processors.DFTypes.DATE
        expected[["duniques"]] = df_processors.DFTypes.UNIQUE
        expected[
            ["dnumerics1", "dnumerics2", "dnumerics3", "dmissing"]
        ] = df_processors.DFTypes.NUMERIC  # fmt: skip
        assert_series_equal(
            self.column_stats[self.columns].loc["types"], expected[self.columns]
        )

    def test_uniques_summary(self):
        expected = pd.Series(
            index=["counts", "uniques", "missing", "missing_perc", "types"],
            data=[self.size, self.size, 0, "0%", df_processors.DFTypes.UNIQUE],
            name="duniques",
            dtype=object,
        )
        assert_series_equal(
            df_processors.get_df_column_summary(self.df, "duniques"), expected
        )

    def test_constant_summary(self):
        self.assertEqual(
            df_processors.get_df_column_summary(self.df, "dconstant"),
            "This is a constant value: a",
        )

    def test_bool1_summary(self):
        count_values = self.df["dbool1"].value_counts()
        total_count = self.df["dbool1"].count()
        count0 = count_values[0]
        count1 = count_values[1]
        perc0 = to_percentage(count0 / total_count)
        perc1 = to_percentage(count1 / total_count)
        expected = pd.Series(
            index=[
                '"0" count',
                '"0" perc',
                '"1" count',
                '"1" perc',
                "counts",
                "uniques",
                "missing",
                "missing_perc",
                "types",
            ],
            data=[
                str(count0),
                perc0,
                str(count1),
                perc1,
                self.size,
                2,
                0,
                "0%",
                df_processors.DFTypes.BOOL,
            ],
            name="dbool1",
            dtype=object,
        )

        assert_series_equal(
            df_processors.get_df_column_summary(self.df, "dbool1"), expected
        )

    def test_bool2_summary(self):
        count_values = self.df["dbool2"].value_counts()
        total_count = self.df["dbool2"].count()
        count0 = count_values["a"]
        count1 = count_values["b"]
        perc0 = to_percentage(count0 / total_count)
        perc1 = to_percentage(count1 / total_count)
        expected = pd.Series(
            index=[
                '"a" count',
                '"a" perc',
                '"b" count',
                '"b" perc',
                "counts",
                "uniques",
                "missing",
                "missing_perc",
                "types",
            ],
            data=[
                str(count0),
                perc0,
                str(count1),
                perc1,
                self.size,
                2,
                0,
                "0%",
                df_processors.DFTypes.BOOL,
            ],
            name="dbool2",
            dtype=object,
        )

        assert_series_equal(
            df_processors.get_df_column_summary(self.df, "dbool2"), expected
        )

    def test_categorical_summary(self):
        expected = pd.Series(
            index=["top", "counts", "uniques", "missing", "missing_perc", "types"],
            data=["a: 500", self.size, 3, 0, "0%", df_processors.DFTypes.CATEGORICAL],
            name="dcategoricals",
            dtype=object,
        )

        assert_series_equal(
            df_processors.get_df_column_summary(self.df, "dcategoricals"), expected
        )

    def test_dates_summary(self):
        dmin = self.df["ddates"].min()
        dmax = self.df["ddates"].max()
        expected = pd.Series(
            index=[
                "max",
                "min",
                "range",
                "counts",
                "uniques",
                "missing",
                "missing_perc",
                "types",
            ],
            data=[
                dmax,
                dmin,
                dmax - dmin,
                self.size,
                self.size,
                0,
                "0%",
                df_processors.DFTypes.DATE,
            ],
            name="ddates",
            dtype=object,
        ).sort_index()

        tmp = df_processors.get_df_column_summary(self.df, "ddates").sort_index()
        assert_series_equal(tmp, expected)

    def test_numerics_summary(self):
        num1 = self.df["dnumerics1"]
        dm, dmp = df_processors.get_deviation_of_mean(num1)
        dam, damp = df_processors.get_median_absolute_deviation(num1)
        expected = pd.Series(
            index=[
                "mean",
                "std",
                "variance",
                "min",
                "max",
                "mode",
                "5%",
                "25%",
                "50%",
                "75%",
                "95%",
                "iqr",
                "kurtosis",
                "skewness",
                "sum",
                "mad",
                "cv",
                "zeros_num",
                "zeros_perc",
                "deviating_of_mean",
                "deviating_of_mean_perc",
                "deviating_of_median",
                "deviating_of_median_perc",
                "counts",
                "uniques",
                "missing",
                "missing_perc",
                "types",
            ],
            data=[
                num1.mean(),
                num1.std(),
                num1.var(),
                num1.min(),
                num1.max(),
                num1.mode()[0],
                num1.quantile(0.05),
                num1.quantile(0.25),
                num1.quantile(0.5),
                num1.quantile(0.75),
                num1.quantile(0.95),
                num1.quantile(0.75) - num1.quantile(0.25),
                num1.kurt(),
                num1.skew(),
                num1.sum(),
                df_processors.mad(num1),
                num1.std() / num1.mean() if num1.mean() else np.nan,
                self.size - np.count_nonzero(num1),
                to_percentage((self.size - np.count_nonzero(num1)) / self.size),
                dm,
                dmp,
                dam,
                damp,
                self.size,
                self.size,
                0,
                "0%",
                df_processors.DFTypes.NUMERIC,
            ],
            name="dnumerics1",
            dtype=object,
        )

        assert_series_equal(
            df_processors.get_df_column_summary(self.df, "dnumerics1"), expected
        )
