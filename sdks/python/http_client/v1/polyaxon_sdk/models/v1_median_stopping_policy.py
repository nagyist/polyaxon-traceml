#!/usr/bin/python
#
# Copyright 2018-2021 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

    Polyaxon SDKs and REST API specification.  # noqa: E501

    The version of the OpenAPI document: 1.7.2
    Contact: contact@polyaxon.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from polyaxon_sdk.configuration import Configuration


class V1MedianStoppingPolicy(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'kind': 'str',
        'evaluation_interval': 'int',
        'min_interval': 'int',
        'min_samples': 'int'
    }

    attribute_map = {
        'kind': 'kind',
        'evaluation_interval': 'evaluationInterval',
        'min_interval': 'minInterval',
        'min_samples': 'minSamples'
    }

    def __init__(self, kind='median', evaluation_interval=None, min_interval=None, min_samples=None, local_vars_configuration=None):  # noqa: E501
        """V1MedianStoppingPolicy - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._kind = None
        self._evaluation_interval = None
        self._min_interval = None
        self._min_samples = None
        self.discriminator = None

        if kind is not None:
            self.kind = kind
        if evaluation_interval is not None:
            self.evaluation_interval = evaluation_interval
        if min_interval is not None:
            self.min_interval = min_interval
        if min_samples is not None:
            self.min_samples = min_samples

    @property
    def kind(self):
        """Gets the kind of this V1MedianStoppingPolicy.  # noqa: E501


        :return: The kind of this V1MedianStoppingPolicy.  # noqa: E501
        :rtype: str
        """
        return self._kind

    @kind.setter
    def kind(self, kind):
        """Sets the kind of this V1MedianStoppingPolicy.


        :param kind: The kind of this V1MedianStoppingPolicy.  # noqa: E501
        :type: str
        """

        self._kind = kind

    @property
    def evaluation_interval(self):
        """Gets the evaluation_interval of this V1MedianStoppingPolicy.  # noqa: E501

        Interval/Frequency for applying the policy.  # noqa: E501

        :return: The evaluation_interval of this V1MedianStoppingPolicy.  # noqa: E501
        :rtype: int
        """
        return self._evaluation_interval

    @evaluation_interval.setter
    def evaluation_interval(self, evaluation_interval):
        """Sets the evaluation_interval of this V1MedianStoppingPolicy.

        Interval/Frequency for applying the policy.  # noqa: E501

        :param evaluation_interval: The evaluation_interval of this V1MedianStoppingPolicy.  # noqa: E501
        :type: int
        """

        self._evaluation_interval = evaluation_interval

    @property
    def min_interval(self):
        """Gets the min_interval of this V1MedianStoppingPolicy.  # noqa: E501


        :return: The min_interval of this V1MedianStoppingPolicy.  # noqa: E501
        :rtype: int
        """
        return self._min_interval

    @min_interval.setter
    def min_interval(self, min_interval):
        """Sets the min_interval of this V1MedianStoppingPolicy.


        :param min_interval: The min_interval of this V1MedianStoppingPolicy.  # noqa: E501
        :type: int
        """

        self._min_interval = min_interval

    @property
    def min_samples(self):
        """Gets the min_samples of this V1MedianStoppingPolicy.  # noqa: E501


        :return: The min_samples of this V1MedianStoppingPolicy.  # noqa: E501
        :rtype: int
        """
        return self._min_samples

    @min_samples.setter
    def min_samples(self, min_samples):
        """Sets the min_samples of this V1MedianStoppingPolicy.


        :param min_samples: The min_samples of this V1MedianStoppingPolicy.  # noqa: E501
        :type: int
        """

        self._min_samples = min_samples

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1MedianStoppingPolicy):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1MedianStoppingPolicy):
            return True

        return self.to_dict() != other.to_dict()
