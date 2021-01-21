// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.5.3
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1ModelVersion from './V1ModelVersion';

/**
 * The V1ListModelVersionsResponse model module.
 * @module model/V1ListModelVersionsResponse
 * @version 1.5.3
 */
class V1ListModelVersionsResponse {
    /**
     * Constructs a new <code>V1ListModelVersionsResponse</code>.
     * @alias module:model/V1ListModelVersionsResponse
     */
    constructor() { 
        
        V1ListModelVersionsResponse.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1ListModelVersionsResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1ListModelVersionsResponse} obj Optional instance to populate.
     * @return {module:model/V1ListModelVersionsResponse} The populated <code>V1ListModelVersionsResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1ListModelVersionsResponse();

            if (data.hasOwnProperty('count')) {
                obj['count'] = ApiClient.convertToType(data['count'], 'Number');
            }
            if (data.hasOwnProperty('results')) {
                obj['results'] = ApiClient.convertToType(data['results'], [V1ModelVersion]);
            }
            if (data.hasOwnProperty('previous')) {
                obj['previous'] = ApiClient.convertToType(data['previous'], 'String');
            }
            if (data.hasOwnProperty('next')) {
                obj['next'] = ApiClient.convertToType(data['next'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} count
 */
V1ListModelVersionsResponse.prototype['count'] = undefined;

/**
 * @member {Array.<module:model/V1ModelVersion>} results
 */
V1ListModelVersionsResponse.prototype['results'] = undefined;

/**
 * @member {String} previous
 */
V1ListModelVersionsResponse.prototype['previous'] = undefined;

/**
 * @member {String} next
 */
V1ListModelVersionsResponse.prototype['next'] = undefined;






export default V1ListModelVersionsResponse;

