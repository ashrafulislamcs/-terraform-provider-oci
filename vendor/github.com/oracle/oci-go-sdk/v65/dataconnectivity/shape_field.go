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

// ShapeField The shape field object.
type ShapeField struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// A detailed description of the object.
	Description *string `mandatory:"false" json:"description"`

	// The type reference.
	Type *interface{} `mandatory:"false" json:"type"`

	// Labels are keywords or labels that you can add to data assets, dataflows, and so on. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`

	NativeShapeField *NativeShapeField `mandatory:"false" json:"nativeShapeField"`
}

//GetKey returns Key
func (m ShapeField) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m ShapeField) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m ShapeField) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetConfigValues returns ConfigValues
func (m ShapeField) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m ShapeField) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetName returns Name
func (m ShapeField) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m ShapeField) GetDescription() *string {
	return m.Description
}

func (m ShapeField) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShapeField) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ShapeField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeShapeField ShapeField
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeShapeField
	}{
		"SHAPE_FIELD",
		(MarshalTypeShapeField)(m),
	}

	return json.Marshal(&s)
}
