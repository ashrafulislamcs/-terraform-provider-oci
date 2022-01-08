// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v54/common"
	"net/http"
	"strings"
)

// ListProblemEndpointsRequest wrapper for the ListProblemEndpoints operation
type ListProblemEndpointsRequest struct {

	// OCId of the problem.
	ProblemId *string `mandatory:"true" contributesTo:"path" name:"problemId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListProblemEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListProblemEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProblemEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProblemEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProblemEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProblemEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProblemEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingListProblemEndpointsSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProblemEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := mappingListProblemEndpointsSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProblemEndpointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProblemEndpointsResponse wrapper for the ListProblemEndpoints operation
type ListProblemEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProblemEndpointCollection instances
	ProblemEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProblemEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProblemEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProblemEndpointsSortOrderEnum Enum with underlying type: string
type ListProblemEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListProblemEndpointsSortOrderEnum
const (
	ListProblemEndpointsSortOrderAsc  ListProblemEndpointsSortOrderEnum = "ASC"
	ListProblemEndpointsSortOrderDesc ListProblemEndpointsSortOrderEnum = "DESC"
)

var mappingListProblemEndpointsSortOrderEnum = map[string]ListProblemEndpointsSortOrderEnum{
	"ASC":  ListProblemEndpointsSortOrderAsc,
	"DESC": ListProblemEndpointsSortOrderDesc,
}

// GetListProblemEndpointsSortOrderEnumValues Enumerates the set of values for ListProblemEndpointsSortOrderEnum
func GetListProblemEndpointsSortOrderEnumValues() []ListProblemEndpointsSortOrderEnum {
	values := make([]ListProblemEndpointsSortOrderEnum, 0)
	for _, v := range mappingListProblemEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProblemEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListProblemEndpointsSortOrderEnum
func GetListProblemEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// ListProblemEndpointsSortByEnum Enum with underlying type: string
type ListProblemEndpointsSortByEnum string

// Set of constants representing the allowable values for ListProblemEndpointsSortByEnum
const (
	ListProblemEndpointsSortByTimecreated ListProblemEndpointsSortByEnum = "timeCreated"
)

var mappingListProblemEndpointsSortByEnum = map[string]ListProblemEndpointsSortByEnum{
	"timeCreated": ListProblemEndpointsSortByTimecreated,
}

// GetListProblemEndpointsSortByEnumValues Enumerates the set of values for ListProblemEndpointsSortByEnum
func GetListProblemEndpointsSortByEnumValues() []ListProblemEndpointsSortByEnum {
	values := make([]ListProblemEndpointsSortByEnum, 0)
	for _, v := range mappingListProblemEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProblemEndpointsSortByEnumStringValues Enumerates the set of values in String for ListProblemEndpointsSortByEnum
func GetListProblemEndpointsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}
