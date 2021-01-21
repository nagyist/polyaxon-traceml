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

/**
 * The V1Organization model module.
 * @module model/V1Organization
 * @version 1.5.3
 */
class V1Organization {
    /**
     * Constructs a new <code>V1Organization</code>.
     * @alias module:model/V1Organization
     */
    constructor() { 
        
        V1Organization.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Organization</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Organization} obj Optional instance to populate.
     * @return {module:model/V1Organization} The populated <code>V1Organization</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Organization();

            if (data.hasOwnProperty('user')) {
                obj['user'] = ApiClient.convertToType(data['user'], 'String');
            }
            if (data.hasOwnProperty('user_email')) {
                obj['user_email'] = ApiClient.convertToType(data['user_email'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('is_public')) {
                obj['is_public'] = ApiClient.convertToType(data['is_public'], 'Boolean');
            }
            if (data.hasOwnProperty('created_at')) {
                obj['created_at'] = ApiClient.convertToType(data['created_at'], 'Date');
            }
            if (data.hasOwnProperty('updated_at')) {
                obj['updated_at'] = ApiClient.convertToType(data['updated_at'], 'Date');
            }
            if (data.hasOwnProperty('role')) {
                obj['role'] = ApiClient.convertToType(data['role'], 'String');
            }
            if (data.hasOwnProperty('queue')) {
                obj['queue'] = ApiClient.convertToType(data['queue'], 'String');
            }
            if (data.hasOwnProperty('preset')) {
                obj['preset'] = ApiClient.convertToType(data['preset'], 'String');
            }
            if (data.hasOwnProperty('auth')) {
                obj['auth'] = ApiClient.convertToType(data['auth'], Object);
            }
            if (data.hasOwnProperty('plan')) {
                obj['plan'] = ApiClient.convertToType(data['plan'], Object);
            }
            if (data.hasOwnProperty('usage')) {
                obj['usage'] = ApiClient.convertToType(data['usage'], Object);
            }
        }
        return obj;
    }


}

/**
 * @member {String} user
 */
V1Organization.prototype['user'] = undefined;

/**
 * @member {String} user_email
 */
V1Organization.prototype['user_email'] = undefined;

/**
 * @member {String} name
 */
V1Organization.prototype['name'] = undefined;

/**
 * @member {Boolean} is_public
 */
V1Organization.prototype['is_public'] = undefined;

/**
 * @member {Date} created_at
 */
V1Organization.prototype['created_at'] = undefined;

/**
 * @member {Date} updated_at
 */
V1Organization.prototype['updated_at'] = undefined;

/**
 * @member {String} role
 */
V1Organization.prototype['role'] = undefined;

/**
 * @member {String} queue
 */
V1Organization.prototype['queue'] = undefined;

/**
 * @member {String} preset
 */
V1Organization.prototype['preset'] = undefined;

/**
 * @member {Object} auth
 */
V1Organization.prototype['auth'] = undefined;

/**
 * @member {Object} plan
 */
V1Organization.prototype['plan'] = undefined;

/**
 * @member {Object} usage
 */
V1Organization.prototype['usage'] = undefined;






export default V1Organization;

