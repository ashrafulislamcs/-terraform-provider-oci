// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Queue API
//
// A description of the Queue API
//

package queue

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateMessagesResultEntry Represents the result of a UpdateMessages request, whether it was successful or not.
// If a message was successfully updated in the queue, the entry includes the `id` and `visibleAfter` fields.
// If a message failed to be updated in the queue, the entry includes the `errorCode` and `errorMessage` fields.
type UpdateMessagesResultEntry struct {

	// The id of the message that's been updated.
	Id *int64 `mandatory:"false" json:"id"`

	// The time after which the message will be visible to other consumers. An RFC3339 formatted datetime string
	VisibleAfter *common.SDKTime `mandatory:"false" json:"visibleAfter"`

	// The error code, in case the message was not successfully updated in the queue.
	ErrorCode *int `mandatory:"false" json:"errorCode"`

	// A human-readable error message associated with the error code.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`
}

func (m UpdateMessagesResultEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMessagesResultEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
