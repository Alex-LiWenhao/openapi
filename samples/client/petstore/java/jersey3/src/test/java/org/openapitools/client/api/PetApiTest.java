/*
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.*;
import org.openapitools.client.auth.*;
import java.io.File;
import org.openapitools.client.model.*;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;

/**
 * API tests for PetApi
 */
public class PetApiTest {

    private final PetApi api = new PetApi();
    private final long petId = 5638l;

    /**
     * Add a new pet to the store
     *
     * 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void addPetTest() throws ApiException {
        // add pet
        Pet body = new Pet();
        body.setId(petId);
        body.setName("jersey2 java8 pet");
        Category category = new Category();
        category.setId(petId);
        category.setName("jersey2 java8 category");
        body.setCategory(category);
        body.setStatus(Pet.StatusEnum.AVAILABLE);
        body.setPhotoUrls(Arrays.asList("A", "B", "C"));
        Tag tag = new Tag();
        tag.setId(petId);
        tag.setName("jersey2 java8 tag");
        body.setTags(Arrays.asList(tag));

        api.addPet(body);

        //get pet by ID
        Pet result = api.getPetById(petId);
        Assertions.assertEquals(result.getId(), body.getId());
        Assertions.assertEquals(result.getCategory(), category);
        Assertions.assertEquals(result.getName(), body.getName());
        Assertions.assertEquals(result.getPhotoUrls(), body.getPhotoUrls());
        Assertions.assertEquals(result.getStatus(), body.getStatus());
        Assertions.assertEquals(result.getTags(), body.getTags());

        // update pet
        api.updatePetWithForm(petId, "jersey2 java8 pet 2", "sold");

        //get pet by ID
        Pet result2 = api.getPetById(petId);
        Assertions.assertEquals(result2.getId(), body.getId());
        Assertions.assertEquals(result2.getCategory(), category);
        Assertions.assertEquals(result2.getName(), "jersey2 java8 pet 2");
        Assertions.assertEquals(result2.getPhotoUrls(), body.getPhotoUrls());
        Assertions.assertEquals(result2.getStatus(), Pet.StatusEnum.SOLD);
        Assertions.assertEquals(result2.getTags(), body.getTags());

        // delete pet
        api.deletePet(petId, "empty api key");

        try {
            Pet result3 = api.getPetById(petId);
            Assertions.assertEquals(false, true);
        } catch (ApiException e) {
//            System.err.println("Exception when calling PetApi#getPetById");
//            System.err.println("Status code: " + e.getCode());
//            System.err.println("Reason: " + e.getResponseBody());
//            System.err.println("Response headers: " + e.getResponseHeaders());

            Assertions.assertEquals(e.getCode(), 404);
            Assertions.assertEquals(e.getResponseBody(), "{\"code\":1,\"type\":\"error\",\"message\":\"Pet not found\"}");

        }
    }

    /**
     * Deletes a pet
     *
     * 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deletePetTest() throws ApiException {
        //Long petId = null;
        //String apiKey = null;
        //api.deletePet(petId, apiKey);
        // TODO: test validations
    }

    /**
     * Finds Pets by status
     *
     * Multiple status values can be provided with comma separated strings
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void findPetsByStatusTest() throws ApiException {
        //List<String> status = null;
        //List<Pet> response = api.findPetsByStatus(status);
        // TODO: test validations
    }

    /**
     * Finds Pets by tags
     *
     * Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void findPetsByTagsTest() throws ApiException {
        //List<String> tags = null;
        //List<Pet> response = api.findPetsByTags(tags);
        // TODO: test validations
    }

    /**
     * Find pet by ID
     *
     * Returns a single pet
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getPetByIdTest() throws ApiException {
        //Long petId = null;
        //Pet response = api.getPetById(petId);
        // TODO: test validations
    }

    /**
     * Update an existing pet
     *
     * 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updatePetTest() throws ApiException {
        //Pet pet = null;
        //api.updatePet(pet);
        // TODO: test validations
    }

    /**
     * Updates a pet in the store with form data
     *
     * 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updatePetWithFormTest() throws ApiException {
        //Long petId = null;
        //String name = null;
        //String status = null;
        //api.updatePetWithForm(petId, name, status);
        // TODO: test validations
    }

    /**
     * uploads an image
     *
     * 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void uploadFileTest() throws ApiException {
        //Long petId = null;
        //String additionalMetadata = null;
        //File _file = null;
        //ModelApiResponse response = api.uploadFile(petId, additionalMetadata, _file);
        // TODO: test validations
    }

    /**
     * uploads an image (required)
     *
     * 
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void uploadFileWithRequiredFileTest() throws ApiException {
        //Long petId = null;
        //File requiredFile = null;
        //String additionalMetadata = null;
        //ModelApiResponse response = api.uploadFileWithRequiredFile(petId, requiredFile, additionalMetadata);
        // TODO: test validations
    }

}
