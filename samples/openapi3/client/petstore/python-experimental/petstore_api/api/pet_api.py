# coding: utf-8

"""
    OpenAPI Petstore

    This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Generated by: https://openapi-generator.tech
"""

from petstore_api.api_client import ApiClient
from petstore_api.api.pet_api_endpoints.add_pet import AddPet
from petstore_api.api.pet_api_endpoints.delete_pet import DeletePet
from petstore_api.api.pet_api_endpoints.find_pets_by_status import FindPetsByStatus
from petstore_api.api.pet_api_endpoints.find_pets_by_tags import FindPetsByTags
from petstore_api.api.pet_api_endpoints.get_pet_by_id import GetPetById
from petstore_api.api.pet_api_endpoints.update_pet import UpdatePet
from petstore_api.api.pet_api_endpoints.update_pet_with_form import UpdatePetWithForm
from petstore_api.api.pet_api_endpoints.upload_file_with_required_file import UploadFileWithRequiredFile
from petstore_api.api.pet_api_endpoints.upload_image import UploadImage


class PetApi(
    AddPet,
    DeletePet,
    FindPetsByStatus,
    FindPetsByTags,
    GetPetById,
    UpdatePet,
    UpdatePetWithForm,
    UploadFileWithRequiredFile,
    UploadImage,
    ApiClient,
):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """
    pass
