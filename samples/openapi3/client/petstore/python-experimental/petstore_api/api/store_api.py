# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""

from petstore_api.api_client import ApiClient
from petstore_api.api.store_api_endpoints.delete_order import DeleteOrder
from petstore_api.api.store_api_endpoints.get_inventory import GetInventory
from petstore_api.api.store_api_endpoints.get_order_by_id import GetOrderById
from petstore_api.api.store_api_endpoints.place_order import PlaceOrder


class StoreApi(
    DeleteOrder,
    GetInventory,
    GetOrderById,
    PlaceOrder,
    ApiClient,
):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """
    pass
