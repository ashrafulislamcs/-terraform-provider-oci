// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v54/common"
	"strings"
)

// InstalledPackageSummary A software package installed on a managed instance
type InstalledPackageSummary struct {

	// Package name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique identifier for the package. NOTE - This is not an OCID
	Name *string `mandatory:"true" json:"name"`

	// Type of the package
	Type *string `mandatory:"true" json:"type"`

	// Version of the installed package
	Version *string `mandatory:"true" json:"version"`

	// The architecture for which this package was built
	Architecture *string `mandatory:"false" json:"architecture"`

	// Install time of the package
	InstallTime *string `mandatory:"false" json:"installTime"`

	// date the package was issued by a providing erratum (if available)
	Issued *string `mandatory:"false" json:"issued"`

	// list of software sources that provide the software package
	SoftwareSources []SoftwareSourceId `mandatory:"false" json:"softwareSources"`
}

func (m InstalledPackageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstalledPackageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
