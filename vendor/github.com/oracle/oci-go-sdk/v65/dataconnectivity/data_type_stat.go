// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataTypeStat Statistical data in profiling results.
type DataTypeStat struct {

	// Value of the confidence of the profile result.
	Value *string `mandatory:"false" json:"value"`

	// Placeholder for now, in future we will return the confidence of the profile result (because we are using sampled data and not whole data)
	Confidence *int `mandatory:"false" json:"confidence"`

	// The number of times the value appeared.
	Freq *int64 `mandatory:"false" json:"freq"`

	// Frequency percentage across the sampled row counts (excluding nulls).
	FreqPercentage *float64 `mandatory:"false" json:"freqPercentage"`
}

func (m DataTypeStat) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataTypeStat) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
