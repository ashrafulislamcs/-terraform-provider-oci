// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

// DataModelFormatEnum Enum with underlying type: string
type DataModelFormatEnum string

// Set of constants representing the allowable values for DataModelFormatEnum
const (
	DataModelFormatXml DataModelFormatEnum = "XML"
)

var mappingDataModelFormatEnum = map[string]DataModelFormatEnum{
	"XML": DataModelFormatXml,
}

// GetDataModelFormatEnumValues Enumerates the set of values for DataModelFormatEnum
func GetDataModelFormatEnumValues() []DataModelFormatEnum {
	values := make([]DataModelFormatEnum, 0)
	for _, v := range mappingDataModelFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetDataModelFormatEnumStringValues Enumerates the set of values in String for DataModelFormatEnum
func GetDataModelFormatEnumStringValues() []string {
	return []string{
		"XML",
	}
}
