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
import V1ArtifactKind from './V1ArtifactKind';

/**
 * The V1RunArtifact model module.
 * @module model/V1RunArtifact
 * @version 1.5.3
 */
class V1RunArtifact {
    /**
     * Constructs a new <code>V1RunArtifact</code>.
     * @alias module:model/V1RunArtifact
     */
    constructor() { 
        
        V1RunArtifact.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1RunArtifact</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1RunArtifact} obj Optional instance to populate.
     * @return {module:model/V1RunArtifact} The populated <code>V1RunArtifact</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1RunArtifact();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('state')) {
                obj['state'] = ApiClient.convertToType(data['state'], 'String');
            }
            if (data.hasOwnProperty('kind')) {
                obj['kind'] = V1ArtifactKind.constructFromObject(data['kind']);
            }
            if (data.hasOwnProperty('path')) {
                obj['path'] = ApiClient.convertToType(data['path'], 'String');
            }
            if (data.hasOwnProperty('connection')) {
                obj['connection'] = ApiClient.convertToType(data['connection'], 'String');
            }
            if (data.hasOwnProperty('summary')) {
                obj['summary'] = ApiClient.convertToType(data['summary'], Object);
            }
            if (data.hasOwnProperty('is_input')) {
                obj['is_input'] = ApiClient.convertToType(data['is_input'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {String} name
 */
V1RunArtifact.prototype['name'] = undefined;

/**
 * @member {String} state
 */
V1RunArtifact.prototype['state'] = undefined;

/**
 * @member {module:model/V1ArtifactKind} kind
 */
V1RunArtifact.prototype['kind'] = undefined;

/**
 * @member {String} path
 */
V1RunArtifact.prototype['path'] = undefined;

/**
 * @member {String} connection
 */
V1RunArtifact.prototype['connection'] = undefined;

/**
 * @member {Object} summary
 */
V1RunArtifact.prototype['summary'] = undefined;

/**
 * @member {Boolean} is_input
 */
V1RunArtifact.prototype['is_input'] = undefined;






export default V1RunArtifact;

