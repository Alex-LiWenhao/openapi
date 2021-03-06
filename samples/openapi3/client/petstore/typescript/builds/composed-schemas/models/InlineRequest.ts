/**
 * Example
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { Cat } from './Cat';
import { Dog } from './Dog';
import { HttpFile } from '../http/http';

export class InlineRequest {
    'hunts'?: boolean;
    'age'?: number;
    'bark'?: boolean;
    'breed'?: InlineRequestBreedEnum;

    static readonly discriminator: string | undefined = "petType";

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "hunts",
            "baseName": "hunts",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "age",
            "baseName": "age",
            "type": "number",
            "format": ""
        },
        {
            "name": "bark",
            "baseName": "bark",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "breed",
            "baseName": "breed",
            "type": "InlineRequestBreedEnum",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return InlineRequest.attributeTypeMap;
    }

    public constructor() {
        this.petType = "InlineRequest";
    }
}


export type InlineRequestBreedEnum = "Dingo" | "Husky" | "Retriever" | "Shepherd" ;

