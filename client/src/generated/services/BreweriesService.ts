/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { controllers_BreweryInput } from '../models/controllers_BreweryInput';
import type { controllers_BreweryResponse } from '../models/controllers_BreweryResponse';
import type { controllers_DeletedBreweryResponse } from '../models/controllers_DeletedBreweryResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class BreweriesService {

    /**
     * Get a list of all breweries
     * Get a list of all breweries
     * @returns controllers_BreweryResponse OK
     * @throws ApiError
     */
    public static getBreweries(): CancelablePromise<Array<controllers_BreweryResponse>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/breweries',
            errors: {
                404: `Not Found`,
            },
        });
    }

    /**
     * Create a brewery
     * @param brewery Brewery input payload
     * @returns controllers_BreweryResponse OK
     * @throws ApiError
     */
    public static postBreweries(
        brewery: controllers_BreweryInput,
    ): CancelablePromise<controllers_BreweryResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/breweries',
            body: brewery,
            errors: {
                400: `Bad Request`,
            },
        });
    }

    /**
     * Get a single brewery by ID
     * @param id Brewery ID
     * @returns controllers_BreweryResponse OK
     * @throws ApiError
     */
    public static getBreweries1(
        id: number,
    ): CancelablePromise<controllers_BreweryResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/breweries/{id}',
            path: {
                'id': id,
            },
            errors: {
                404: `Not Found`,
            },
        });
    }

    /**
     * Delete a brewery
     * Delete a brewery only if there are no beers associated with it
     * @param id Brewery ID
     * @returns controllers_DeletedBreweryResponse OK
     * @throws ApiError
     */
    public static deleteBreweries(
        id: number,
    ): CancelablePromise<controllers_DeletedBreweryResponse> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/breweries/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
            },
        });
    }

    /**
     * Update a brewery
     * Update a brewery by id
     * @param id Brewery ID
     * @param brewery Brewery Payload
     * @returns controllers_BreweryResponse OK
     * @throws ApiError
     */
    public static patchBreweries(
        id: number,
        brewery: controllers_BreweryInput,
    ): CancelablePromise<controllers_BreweryResponse> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/breweries/{id}',
            path: {
                'id': id,
            },
            body: brewery,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
            },
        });
    }

}
