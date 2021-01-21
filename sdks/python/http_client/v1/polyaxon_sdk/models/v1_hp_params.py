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

    The version of the OpenAPI document: 1.5.3
    Contact: contact@polyaxon.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from polyaxon_sdk.configuration import Configuration


class V1HpParams(object):
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
        'choice': 'V1HpChoice',
        'pchoice': 'V1HpPChoice',
        'range': 'V1HpRange',
        'linspace': 'V1HpLinSpace',
        'logspace': 'V1HpLogSpace',
        'geomspace': 'V1HpGeomSpace',
        'uniform': 'V1HpUniform',
        'quniform': 'V1HpQUniform',
        'loguniform': 'V1HpLogUniform',
        'qloguniform': 'V1HpQLogUniform',
        'normal': 'V1HpNormal',
        'qnormal': 'V1HpQNormal',
        'lognormal': 'V1HpLogNormal',
        'qlognormal': 'V1HpQLogNormal'
    }

    attribute_map = {
        'choice': 'choice',
        'pchoice': 'pchoice',
        'range': 'range',
        'linspace': 'linspace',
        'logspace': 'logspace',
        'geomspace': 'geomspace',
        'uniform': 'uniform',
        'quniform': 'quniform',
        'loguniform': 'loguniform',
        'qloguniform': 'qloguniform',
        'normal': 'normal',
        'qnormal': 'qnormal',
        'lognormal': 'lognormal',
        'qlognormal': 'qlognormal'
    }

    def __init__(self, choice=None, pchoice=None, range=None, linspace=None, logspace=None, geomspace=None, uniform=None, quniform=None, loguniform=None, qloguniform=None, normal=None, qnormal=None, lognormal=None, qlognormal=None, local_vars_configuration=None):  # noqa: E501
        """V1HpParams - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._choice = None
        self._pchoice = None
        self._range = None
        self._linspace = None
        self._logspace = None
        self._geomspace = None
        self._uniform = None
        self._quniform = None
        self._loguniform = None
        self._qloguniform = None
        self._normal = None
        self._qnormal = None
        self._lognormal = None
        self._qlognormal = None
        self.discriminator = None

        if choice is not None:
            self.choice = choice
        if pchoice is not None:
            self.pchoice = pchoice
        if range is not None:
            self.range = range
        if linspace is not None:
            self.linspace = linspace
        if logspace is not None:
            self.logspace = logspace
        if geomspace is not None:
            self.geomspace = geomspace
        if uniform is not None:
            self.uniform = uniform
        if quniform is not None:
            self.quniform = quniform
        if loguniform is not None:
            self.loguniform = loguniform
        if qloguniform is not None:
            self.qloguniform = qloguniform
        if normal is not None:
            self.normal = normal
        if qnormal is not None:
            self.qnormal = qnormal
        if lognormal is not None:
            self.lognormal = lognormal
        if qlognormal is not None:
            self.qlognormal = qlognormal

    @property
    def choice(self):
        """Gets the choice of this V1HpParams.  # noqa: E501


        :return: The choice of this V1HpParams.  # noqa: E501
        :rtype: V1HpChoice
        """
        return self._choice

    @choice.setter
    def choice(self, choice):
        """Sets the choice of this V1HpParams.


        :param choice: The choice of this V1HpParams.  # noqa: E501
        :type: V1HpChoice
        """

        self._choice = choice

    @property
    def pchoice(self):
        """Gets the pchoice of this V1HpParams.  # noqa: E501


        :return: The pchoice of this V1HpParams.  # noqa: E501
        :rtype: V1HpPChoice
        """
        return self._pchoice

    @pchoice.setter
    def pchoice(self, pchoice):
        """Sets the pchoice of this V1HpParams.


        :param pchoice: The pchoice of this V1HpParams.  # noqa: E501
        :type: V1HpPChoice
        """

        self._pchoice = pchoice

    @property
    def range(self):
        """Gets the range of this V1HpParams.  # noqa: E501


        :return: The range of this V1HpParams.  # noqa: E501
        :rtype: V1HpRange
        """
        return self._range

    @range.setter
    def range(self, range):
        """Sets the range of this V1HpParams.


        :param range: The range of this V1HpParams.  # noqa: E501
        :type: V1HpRange
        """

        self._range = range

    @property
    def linspace(self):
        """Gets the linspace of this V1HpParams.  # noqa: E501


        :return: The linspace of this V1HpParams.  # noqa: E501
        :rtype: V1HpLinSpace
        """
        return self._linspace

    @linspace.setter
    def linspace(self, linspace):
        """Sets the linspace of this V1HpParams.


        :param linspace: The linspace of this V1HpParams.  # noqa: E501
        :type: V1HpLinSpace
        """

        self._linspace = linspace

    @property
    def logspace(self):
        """Gets the logspace of this V1HpParams.  # noqa: E501


        :return: The logspace of this V1HpParams.  # noqa: E501
        :rtype: V1HpLogSpace
        """
        return self._logspace

    @logspace.setter
    def logspace(self, logspace):
        """Sets the logspace of this V1HpParams.


        :param logspace: The logspace of this V1HpParams.  # noqa: E501
        :type: V1HpLogSpace
        """

        self._logspace = logspace

    @property
    def geomspace(self):
        """Gets the geomspace of this V1HpParams.  # noqa: E501


        :return: The geomspace of this V1HpParams.  # noqa: E501
        :rtype: V1HpGeomSpace
        """
        return self._geomspace

    @geomspace.setter
    def geomspace(self, geomspace):
        """Sets the geomspace of this V1HpParams.


        :param geomspace: The geomspace of this V1HpParams.  # noqa: E501
        :type: V1HpGeomSpace
        """

        self._geomspace = geomspace

    @property
    def uniform(self):
        """Gets the uniform of this V1HpParams.  # noqa: E501


        :return: The uniform of this V1HpParams.  # noqa: E501
        :rtype: V1HpUniform
        """
        return self._uniform

    @uniform.setter
    def uniform(self, uniform):
        """Sets the uniform of this V1HpParams.


        :param uniform: The uniform of this V1HpParams.  # noqa: E501
        :type: V1HpUniform
        """

        self._uniform = uniform

    @property
    def quniform(self):
        """Gets the quniform of this V1HpParams.  # noqa: E501


        :return: The quniform of this V1HpParams.  # noqa: E501
        :rtype: V1HpQUniform
        """
        return self._quniform

    @quniform.setter
    def quniform(self, quniform):
        """Sets the quniform of this V1HpParams.


        :param quniform: The quniform of this V1HpParams.  # noqa: E501
        :type: V1HpQUniform
        """

        self._quniform = quniform

    @property
    def loguniform(self):
        """Gets the loguniform of this V1HpParams.  # noqa: E501


        :return: The loguniform of this V1HpParams.  # noqa: E501
        :rtype: V1HpLogUniform
        """
        return self._loguniform

    @loguniform.setter
    def loguniform(self, loguniform):
        """Sets the loguniform of this V1HpParams.


        :param loguniform: The loguniform of this V1HpParams.  # noqa: E501
        :type: V1HpLogUniform
        """

        self._loguniform = loguniform

    @property
    def qloguniform(self):
        """Gets the qloguniform of this V1HpParams.  # noqa: E501


        :return: The qloguniform of this V1HpParams.  # noqa: E501
        :rtype: V1HpQLogUniform
        """
        return self._qloguniform

    @qloguniform.setter
    def qloguniform(self, qloguniform):
        """Sets the qloguniform of this V1HpParams.


        :param qloguniform: The qloguniform of this V1HpParams.  # noqa: E501
        :type: V1HpQLogUniform
        """

        self._qloguniform = qloguniform

    @property
    def normal(self):
        """Gets the normal of this V1HpParams.  # noqa: E501


        :return: The normal of this V1HpParams.  # noqa: E501
        :rtype: V1HpNormal
        """
        return self._normal

    @normal.setter
    def normal(self, normal):
        """Sets the normal of this V1HpParams.


        :param normal: The normal of this V1HpParams.  # noqa: E501
        :type: V1HpNormal
        """

        self._normal = normal

    @property
    def qnormal(self):
        """Gets the qnormal of this V1HpParams.  # noqa: E501


        :return: The qnormal of this V1HpParams.  # noqa: E501
        :rtype: V1HpQNormal
        """
        return self._qnormal

    @qnormal.setter
    def qnormal(self, qnormal):
        """Sets the qnormal of this V1HpParams.


        :param qnormal: The qnormal of this V1HpParams.  # noqa: E501
        :type: V1HpQNormal
        """

        self._qnormal = qnormal

    @property
    def lognormal(self):
        """Gets the lognormal of this V1HpParams.  # noqa: E501


        :return: The lognormal of this V1HpParams.  # noqa: E501
        :rtype: V1HpLogNormal
        """
        return self._lognormal

    @lognormal.setter
    def lognormal(self, lognormal):
        """Sets the lognormal of this V1HpParams.


        :param lognormal: The lognormal of this V1HpParams.  # noqa: E501
        :type: V1HpLogNormal
        """

        self._lognormal = lognormal

    @property
    def qlognormal(self):
        """Gets the qlognormal of this V1HpParams.  # noqa: E501


        :return: The qlognormal of this V1HpParams.  # noqa: E501
        :rtype: V1HpQLogNormal
        """
        return self._qlognormal

    @qlognormal.setter
    def qlognormal(self, qlognormal):
        """Sets the qlognormal of this V1HpParams.


        :param qlognormal: The qlognormal of this V1HpParams.  # noqa: E501
        :type: V1HpQLogNormal
        """

        self._qlognormal = qlognormal

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
        if not isinstance(other, V1HpParams):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1HpParams):
            return True

        return self.to_dict() != other.to_dict()
