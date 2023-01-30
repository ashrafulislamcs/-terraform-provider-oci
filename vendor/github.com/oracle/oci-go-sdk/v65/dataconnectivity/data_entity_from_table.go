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

// DataEntityFromTable The table entity data entity.
type DataEntityFromTable struct {

	// Map<String, String> for entity properties
	EntityProperties map[string]string `mandatory:"false" json:"entityProperties"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The model version of the object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description of the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The external key of the object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	Shape *Shape `mandatory:"false" json:"shape"`

	// The shape ID.
	ShapeId *string `mandatory:"false" json:"shapeId"`

	// Specifies other type label.
	OtherTypeLabel *string `mandatory:"false" json:"otherTypeLabel"`

	// An array of unique keys.
	UniqueKeys []UniqueKey `mandatory:"false" json:"uniqueKeys"`

	// An array of foreign keys.
	ForeignKeys []ForeignKey `mandatory:"false" json:"foreignKeys"`

	// The resource name.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with an upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The entity type.
	EntityType DataEntityFromTableEntityTypeEnum `mandatory:"false" json:"entityType,omitempty"`
}

//GetEntityProperties returns EntityProperties
func (m DataEntityFromTable) GetEntityProperties() map[string]string {
	return m.EntityProperties
}

//GetMetadata returns Metadata
func (m DataEntityFromTable) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m DataEntityFromTable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataEntityFromTable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataEntityFromTableEntityTypeEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetDataEntityFromTableEntityTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataEntityFromTable) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataEntityFromTable DataEntityFromTable
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDataEntityFromTable
	}{
		"TABLE_ENTITY",
		(MarshalTypeDataEntityFromTable)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DataEntityFromTable) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		EntityProperties map[string]string                 `json:"entityProperties"`
		Metadata         *ObjectMetadata                   `json:"metadata"`
		Key              *string                           `json:"key"`
		ModelVersion     *string                           `json:"modelVersion"`
		ParentRef        *ParentReference                  `json:"parentRef"`
		Name             *string                           `json:"name"`
		Description      *string                           `json:"description"`
		ObjectVersion    *int                              `json:"objectVersion"`
		ExternalKey      *string                           `json:"externalKey"`
		Shape            *Shape                            `json:"shape"`
		ShapeId          *string                           `json:"shapeId"`
		EntityType       DataEntityFromTableEntityTypeEnum `json:"entityType"`
		OtherTypeLabel   *string                           `json:"otherTypeLabel"`
		UniqueKeys       []uniquekey                       `json:"uniqueKeys"`
		ForeignKeys      []ForeignKey                      `json:"foreignKeys"`
		ResourceName     *string                           `json:"resourceName"`
		ObjectStatus     *int                              `json:"objectStatus"`
		Identifier       *string                           `json:"identifier"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.EntityProperties = model.EntityProperties

	m.Metadata = model.Metadata

	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.Description = model.Description

	m.ObjectVersion = model.ObjectVersion

	m.ExternalKey = model.ExternalKey

	m.Shape = model.Shape

	m.ShapeId = model.ShapeId

	m.EntityType = model.EntityType

	m.OtherTypeLabel = model.OtherTypeLabel

	m.UniqueKeys = make([]UniqueKey, len(model.UniqueKeys))
	for i, n := range model.UniqueKeys {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.UniqueKeys[i] = nn.(UniqueKey)
		} else {
			m.UniqueKeys[i] = nil
		}
	}

	m.ForeignKeys = make([]ForeignKey, len(model.ForeignKeys))
	for i, n := range model.ForeignKeys {
		m.ForeignKeys[i] = n
	}

	m.ResourceName = model.ResourceName

	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier

	return
}

// DataEntityFromTableEntityTypeEnum Enum with underlying type: string
type DataEntityFromTableEntityTypeEnum string

// Set of constants representing the allowable values for DataEntityFromTableEntityTypeEnum
const (
	DataEntityFromTableEntityTypeTable     DataEntityFromTableEntityTypeEnum = "TABLE"
	DataEntityFromTableEntityTypeView      DataEntityFromTableEntityTypeEnum = "VIEW"
	DataEntityFromTableEntityTypeFile      DataEntityFromTableEntityTypeEnum = "FILE"
	DataEntityFromTableEntityTypeSql       DataEntityFromTableEntityTypeEnum = "SQL"
	DataEntityFromTableEntityTypeDataStore DataEntityFromTableEntityTypeEnum = "DATA_STORE"
	DataEntityFromTableEntityTypeMessage   DataEntityFromTableEntityTypeEnum = "MESSAGE"
)

var mappingDataEntityFromTableEntityTypeEnum = map[string]DataEntityFromTableEntityTypeEnum{
	"TABLE":      DataEntityFromTableEntityTypeTable,
	"VIEW":       DataEntityFromTableEntityTypeView,
	"FILE":       DataEntityFromTableEntityTypeFile,
	"SQL":        DataEntityFromTableEntityTypeSql,
	"DATA_STORE": DataEntityFromTableEntityTypeDataStore,
	"MESSAGE":    DataEntityFromTableEntityTypeMessage,
}

var mappingDataEntityFromTableEntityTypeEnumLowerCase = map[string]DataEntityFromTableEntityTypeEnum{
	"table":      DataEntityFromTableEntityTypeTable,
	"view":       DataEntityFromTableEntityTypeView,
	"file":       DataEntityFromTableEntityTypeFile,
	"sql":        DataEntityFromTableEntityTypeSql,
	"data_store": DataEntityFromTableEntityTypeDataStore,
	"message":    DataEntityFromTableEntityTypeMessage,
}

// GetDataEntityFromTableEntityTypeEnumValues Enumerates the set of values for DataEntityFromTableEntityTypeEnum
func GetDataEntityFromTableEntityTypeEnumValues() []DataEntityFromTableEntityTypeEnum {
	values := make([]DataEntityFromTableEntityTypeEnum, 0)
	for _, v := range mappingDataEntityFromTableEntityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDataEntityFromTableEntityTypeEnumStringValues Enumerates the set of values in String for DataEntityFromTableEntityTypeEnum
func GetDataEntityFromTableEntityTypeEnumStringValues() []string {
	return []string{
		"TABLE",
		"VIEW",
		"FILE",
		"SQL",
		"DATA_STORE",
		"MESSAGE",
	}
}

// GetMappingDataEntityFromTableEntityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataEntityFromTableEntityTypeEnum(val string) (DataEntityFromTableEntityTypeEnum, bool) {
	enum, ok := mappingDataEntityFromTableEntityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
