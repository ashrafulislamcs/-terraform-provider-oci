// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// UsageApi API
//
// A description of the UsageApi API.
//

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v54/common"
	"strings"
)

// RedeemableUser The Email object for redeemable user.
type RedeemableUser struct {

	// The Email Id for redeemable user.
	EmailId *string `mandatory:"true" json:"emailId"`
}

func (m RedeemableUser) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RedeemableUser) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
