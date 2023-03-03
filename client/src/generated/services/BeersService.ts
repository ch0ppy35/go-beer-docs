/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { controllers_BeerInput } from '../models/controllers_BeerInput';
import type { controllers_BeerResponse } from '../models/controllers_BeerResponse';
import type { controllers_DeletedBeerResponse } from '../models/controllers_DeletedBeerResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class BeersService {

    /**
     * Get a list of all beers
     * Get a list of all beers
     * @returns controllers_BeerResponse Successful operation
     * @throws ApiError
     */
    public static getBeers(): CancelablePromise<Array<controllers_BeerResponse>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/beers/',
            errors: {
                404: `Not Found`,
            },
        });
    }

    /**
     * Create a new beer
     * Create a new beer
     * @param input Beer input
     * @returns controllers_BeerResponse Successful operation
     * @throws ApiError
     */
    public static postBeers(
        input: controllers_BeerInput,
    ): CancelablePromise<controllers_BeerResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/beers/',
            body: input,
            errors: {
                400: `Ensure input is correct!`,
            },
        });
    }

    /**
     * Get a beer by ID
     * @param id Beer ID
     * @returns controllers_BeerResponse OK
     * @throws ApiError
     */
    public static getBeers1(
        id: number,
    ): CancelablePromise<controllers_BeerResponse> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/beers/{id}',
            path: {
                'id': id,
            },
            errors: {
                404: `Not Found`,
            },
        });
    }

    /**
     * Delete a beer by ID
     * @param id Beer ID
     * @returns controllers_DeletedBeerResponse OK
     * @throws ApiError
     */
    public static deleteBeers(
        id: number,
    ): CancelablePromise<controllers_DeletedBeerResponse> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/beers/{id}',
            path: {
                'id': id,
            },
            errors: {
                400: `Bad Request`,
            },
        });
    }

    /**
     * Update a beer by ID
     * @param id Beer ID
     * @param beer Beer input payload
     * @returns controllers_BeerResponse OK
     * @throws ApiError
     */
    public static patchBeers(
        id: number,
        beer: controllers_BeerInput,
    ): CancelablePromise<controllers_BeerResponse> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/beers/{id}',
            path: {
                'id': id,
            },
            body: beer,
            errors: {
                400: `Bad Request`,
                404: `Not Found`,
            },
        });
    }

}
