# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from src.swagger_models.base_model_ import Model
from src import util


class DatasetEvaluationElementPoint3d(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """
    def __init__(self, x: float=None, y: float=None, z: float=None):  # noqa: E501
        """DatasetEvaluationElementPoint3d - a model defined in Swagger

        :param x: The x of this DatasetEvaluationElementPoint3d.  # noqa: E501
        :type x: float
        :param y: The y of this DatasetEvaluationElementPoint3d.  # noqa: E501
        :type y: float
        :param z: The z of this DatasetEvaluationElementPoint3d.  # noqa: E501
        :type z: float
        """
        self.swagger_types = {
            'x': float,
            'y': float,
            'z': float
        }

        self.attribute_map = {
            'x': 'x',
            'y': 'y',
            'z': 'z'
        }
        self._x = x
        self._y = y
        self._z = z

    @classmethod
    def from_dict(cls, dikt) -> 'DatasetEvaluationElementPoint3d':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The DatasetEvaluationElementPoint3d of this DatasetEvaluationElementPoint3d.  # noqa: E501
        :rtype: DatasetEvaluationElementPoint3d
        """
        return util.deserialize_model(dikt, cls)

    @property
    def x(self) -> float:
        """Gets the x of this DatasetEvaluationElementPoint3d.


        :return: The x of this DatasetEvaluationElementPoint3d.
        :rtype: float
        """
        return self._x

    @x.setter
    def x(self, x: float):
        """Sets the x of this DatasetEvaluationElementPoint3d.


        :param x: The x of this DatasetEvaluationElementPoint3d.
        :type x: float
        """

        self._x = x

    @property
    def y(self) -> float:
        """Gets the y of this DatasetEvaluationElementPoint3d.


        :return: The y of this DatasetEvaluationElementPoint3d.
        :rtype: float
        """
        return self._y

    @y.setter
    def y(self, y: float):
        """Sets the y of this DatasetEvaluationElementPoint3d.


        :param y: The y of this DatasetEvaluationElementPoint3d.
        :type y: float
        """

        self._y = y

    @property
    def z(self) -> float:
        """Gets the z of this DatasetEvaluationElementPoint3d.


        :return: The z of this DatasetEvaluationElementPoint3d.
        :rtype: float
        """
        return self._z

    @z.setter
    def z(self, z: float):
        """Sets the z of this DatasetEvaluationElementPoint3d.


        :param z: The z of this DatasetEvaluationElementPoint3d.
        :type z: float
        """

        self._z = z