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

// OperationFromApi The API operation object.
type OperationFromApi struct {

	// The operation name. This value is unique.
	Name *string `mandatory:"true" json:"name"`

	// The resource name.
	ResourceName *string `mandatory:"true" json:"resourceName"`

	OperationAttributes AbstractOperationAttributes `mandatory:"false" json:"operationAttributes"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The model version of the object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	Shape *Shape `mandatory:"false" json:"shape"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The external key for the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// The status of an object that can be set to value 1 for shallow reference across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

//GetOperationAttributes returns OperationAttributes
func (m OperationFromApi) GetOperationAttributes() AbstractOperationAttributes {
	return m.OperationAttributes
}

//GetMetadata returns Metadata
func (m OperationFromApi) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m OperationFromApi) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OperationFromApi) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OperationFromApi) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOperationFromApi OperationFromApi
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOperationFromApi
	}{
		"API",
		(MarshalTypeOperationFromApi)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OperationFromApi) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		OperationAttributes abstractoperationattributes `json:"operationAttributes"`
		Metadata            *ObjectMetadata             `json:"metadata"`
		Key                 *string                     `json:"key"`
		ModelVersion        *string                     `json:"modelVersion"`
		ParentRef           *ParentReference            `json:"parentRef"`
		Shape               *Shape                      `json:"shape"`
		ObjectVersion       *int                        `json:"objectVersion"`
		ExternalKey         *string                     `json:"externalKey"`
		ObjectStatus        *int                        `json:"objectStatus"`
		Name                *string                     `json:"name"`
		ResourceName        *string                     `json:"resourceName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.OperationAttributes.UnmarshalPolymorphicJSON(model.OperationAttributes.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.OperationAttributes = nn.(AbstractOperationAttributes)
	} else {
		m.OperationAttributes = nil
	}

	m.Metadata = model.Metadata

	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Shape = model.Shape

	m.ObjectVersion = model.ObjectVersion

	m.ExternalKey = model.ExternalKey

	m.ObjectStatus = model.ObjectStatus

	m.Name = model.Name

	m.ResourceName = model.ResourceName

	return
}
