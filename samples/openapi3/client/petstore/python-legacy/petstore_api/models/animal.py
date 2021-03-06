# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""


try:
    from inspect import getfullargspec
except ImportError:
    from inspect import getargspec as getfullargspec
import pprint
import re  # noqa: F401
import six

from petstore_api.configuration import Configuration


class Animal(object):
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
        'class_name': 'str',
        'color': 'str'
    }

    attribute_map = {
        'class_name': 'className',
        'color': 'color'
    }

    discriminator_value_class_map = {
        'Cat': 'Cat',
        'Dog': 'Dog'
    }

    def __init__(self, class_name=None, color='red', local_vars_configuration=None):  # noqa: E501
        """Animal - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration.get_default_copy()
        self.local_vars_configuration = local_vars_configuration

        self._class_name = None
        self._color = None
        self.discriminator = 'class_name'

        self.class_name = class_name
        if color is not None:
            self.color = color

    @property
    def class_name(self):
        """Gets the class_name of this Animal.  # noqa: E501


        :return: The class_name of this Animal.  # noqa: E501
        :rtype: str
        """
        return self._class_name

    @class_name.setter
    def class_name(self, class_name):
        """Sets the class_name of this Animal.


        :param class_name: The class_name of this Animal.  # noqa: E501
        :type class_name: str
        """
        if self.local_vars_configuration.client_side_validation and class_name is None:  # noqa: E501
            raise ValueError("Invalid value for `class_name`, must not be `None`")  # noqa: E501

        self._class_name = class_name

    @property
    def color(self):
        """Gets the color of this Animal.  # noqa: E501


        :return: The color of this Animal.  # noqa: E501
        :rtype: str
        """
        return self._color

    @color.setter
    def color(self, color):
        """Sets the color of this Animal.


        :param color: The color of this Animal.  # noqa: E501
        :type color: str
        """

        self._color = color

    def get_real_child_model(self, data):
        """Returns the real base class specified by the discriminator"""
        discriminator_key = self.attribute_map[self.discriminator]
        discriminator_value = data[discriminator_key]
        return self.discriminator_value_class_map.get(discriminator_value)

    def to_dict(self, serialize=False):
        """Returns the model properties as a dict"""
        result = {}

        def convert(x):
            if hasattr(x, "to_dict"):
                args = getfullargspec(x.to_dict).args
                if len(args) == 1:
                    return x.to_dict()
                else:
                    return x.to_dict(serialize)
            else:
                return x

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            attr = self.attribute_map.get(attr, attr) if serialize else attr
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: convert(x),
                    value
                ))
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], convert(item[1])),
                    value.items()
                ))
            else:
                result[attr] = convert(value)

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, Animal):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, Animal):
            return True

        return self.to_dict() != other.to_dict()
