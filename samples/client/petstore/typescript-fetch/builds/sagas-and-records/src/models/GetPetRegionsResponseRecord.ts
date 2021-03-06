/* tslint:disable */
/* eslint-disable */
/**
 * OpenAPI Petstore
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import {ApiRecordUtils, knownRecordFactories, appFromJS, NormalizedRecordEntities} from "../runtimeSagasAndRecords";
import {getApiEntitiesState} from "../ApiEntitiesSelectors"
import {List, Record, RecordOf, Map} from 'immutable';
import {Schema, schema, NormalizedSchema} from "normalizr";
import {select, call} from "redux-saga/effects";

import {
    GetPetRegionsResponse,
} from './GetPetRegionsResponse';

import {
    ResponseMeta,
} from './ResponseMeta';

import {
    ResponseMetaRecord,
    responseMetaRecordUtils
} from './ResponseMetaRecord';

export const GetPetRegionsResponseRecordProps = {
    recType: "GetPetRegionsResponseApiRecord" as "GetPetRegionsResponseApiRecord",
    meta: ResponseMetaRecord(),
    data: null as List<List<string | null>> | null,
};

export type GetPetRegionsResponseRecordPropsType = typeof GetPetRegionsResponseRecordProps;
export const GetPetRegionsResponseRecord = Record(GetPetRegionsResponseRecordProps, GetPetRegionsResponseRecordProps.recType);
export type GetPetRegionsResponseRecord = RecordOf<GetPetRegionsResponseRecordPropsType>;

knownRecordFactories.set(GetPetRegionsResponseRecordProps.recType, GetPetRegionsResponseRecord);


class GetPetRegionsResponseRecordUtils extends ApiRecordUtils<GetPetRegionsResponse, GetPetRegionsResponseRecord> {
    public normalize(apiObject: GetPetRegionsResponse, asEntity?: boolean): GetPetRegionsResponse {
        (apiObject as any).recType = GetPetRegionsResponseRecordProps.recType;
        responseMetaRecordUtils.normalize(apiObject.meta);
        if (apiObject.data) { (apiObject as any).data = apiObject.data.map(item => item.map(item2 => item2?.toString())); } 
        return apiObject;
    }

    public toApi(record: GetPetRegionsResponseRecord): GetPetRegionsResponse {
        const apiObject = super.toApi(record);
        apiObject.meta = responseMetaRecordUtils.toApi(record.meta);
        if (record.data) { apiObject.data = record.data.map(item => item.toArray().map(item2 => (item2 ? parseFloat(item2) : null) as number)).toArray(); } 
        return apiObject;
    }

    public fromApiPassthrough(apiObject: GetPetRegionsResponse): List<List<string | null>> {
        return appFromJS(apiObject.data);
    }

    public fromApiPassthroughAsEntities(apiObject: GetPetRegionsResponse): NormalizedRecordEntities {
        console.log("entities revival not supported on this response");
        return {entities: {}, result: List<string>()};
    }
}

export const getPetRegionsResponseRecordUtils = new GetPetRegionsResponseRecordUtils();


