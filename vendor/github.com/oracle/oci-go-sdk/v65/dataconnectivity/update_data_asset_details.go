// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDataAssetDetails Properties used in the update data asset operations.
type UpdateDataAssetDetails struct {

	// All the properties of the data asset in a key-value map format.
	Properties map[string]interface{} `mandatory:"true" json:"properties"`

	// Specific DataAsset Type
	Type *string `mandatory:"true" json:"type"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// User-defined description of the data asset.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with an upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The external key of the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// Additional properties of the data asset.
	AssetProperties map[string]string `mandatory:"false" json:"assetProperties"`

	NativeTypeSystem *TypeSystem `mandatory:"false" json:"nativeTypeSystem"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	DefaultConnection *Connection `mandatory:"false" json:"defaultConnection"`

	// The list of endpoints with which this data asset is associated.
	EndPoints []DpEndpoint `mandatory:"false" json:"endPoints"`
}

func (m UpdateDataAssetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDataAssetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDataAssetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ModelVersion      *string                `json:"modelVersion"`
		ModelType         *string                `json:"modelType"`
		Name              *string                `json:"name"`
		Description       *string                `json:"description"`
		ObjectStatus      *int                   `json:"objectStatus"`
		ObjectVersion     *int                   `json:"objectVersion"`
		Identifier        *string                `json:"identifier"`
		ExternalKey       *string                `json:"externalKey"`
		AssetProperties   map[string]string      `json:"assetProperties"`
		NativeTypeSystem  *TypeSystem            `json:"nativeTypeSystem"`
		RegistryMetadata  *RegistryMetadata      `json:"registryMetadata"`
		Metadata          *ObjectMetadata        `json:"metadata"`
		DefaultConnection *Connection            `json:"defaultConnection"`
		EndPoints         []dpendpoint           `json:"endPoints"`
		Properties        map[string]interface{} `json:"properties"`
		Type              *string                `json:"type"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ModelVersion = model.ModelVersion

	m.ModelType = model.ModelType

	m.Name = model.Name

	m.Description = model.Description

	m.ObjectStatus = model.ObjectStatus

	m.ObjectVersion = model.ObjectVersion

	m.Identifier = model.Identifier

	m.ExternalKey = model.ExternalKey

	m.AssetProperties = model.AssetProperties

	m.NativeTypeSystem = model.NativeTypeSystem

	m.RegistryMetadata = model.RegistryMetadata

	m.Metadata = model.Metadata

	m.DefaultConnection = model.DefaultConnection

	m.EndPoints = make([]DpEndpoint, len(model.EndPoints))
	for i, n := range model.EndPoints {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.EndPoints[i] = nn.(DpEndpoint)
		} else {
			m.EndPoints[i] = nil
		}
	}

	m.Properties = model.Properties

	m.Type = model.Type

	return
}
