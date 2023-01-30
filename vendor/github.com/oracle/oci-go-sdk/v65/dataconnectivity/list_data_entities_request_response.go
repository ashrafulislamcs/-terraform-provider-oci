// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDataEntitiesRequest wrapper for the ListDataEntities operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataconnectivity/ListDataEntities.go.html to see an example of how to use ListDataEntitiesRequest.
type ListDataEntitiesRequest struct {

	// The registry OCID.
	RegistryId *string `mandatory:"true" contributesTo:"path" name:"registryId"`

	// The connection key.
	ConnectionKey *string `mandatory:"true" contributesTo:"path" name:"connectionKey"`

	// The schema resource name used for retrieving schemas.
	SchemaResourceName *string `mandatory:"true" contributesTo:"path" name:"schemaResourceName"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Type of the object to filter the results with.
	Type *string `mandatory:"false" contributesTo:"query" name:"type"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order are by relevance score in descending order).
	SortBy ListDataEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListDataEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This parameter can be used to set the api response type to preview.
	ApiMode ListDataEntitiesApiModeEnum `mandatory:"false" contributesTo:"query" name:"apiMode" omitEmpty:"true"`

	// Used to filter by the name of the object.
	NameList []string `contributesTo:"query" name:"nameList" collectionFormat:"multi"`

	// This parameter can be used to specify whether entity search type is a pattern search.
	IsPattern *bool `mandatory:"false" contributesTo:"query" name:"isPattern"`

	// Endpoint ID used for getDataAssetFullDetails.
	EndpointId *string `mandatory:"false" contributesTo:"query" name:"endpointId"`

	// Artifact type which needs to be listed while listing Artifacts.
	IncludeTypes []string `contributesTo:"query" name:"includeTypes" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataEntitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataEntitiesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataEntitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataEntitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataEntitiesApiModeEnum(string(request.ApiMode)); !ok && request.ApiMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ApiMode: %s. Supported values are: %s.", request.ApiMode, strings.Join(GetListDataEntitiesApiModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataEntitiesResponse wrapper for the ListDataEntities operation
type ListDataEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataEntitySummaryCollection instances
	DataEntitySummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListDataEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataEntitiesSortByEnum Enum with underlying type: string
type ListDataEntitiesSortByEnum string

// Set of constants representing the allowable values for ListDataEntitiesSortByEnum
const (
	ListDataEntitiesSortById          ListDataEntitiesSortByEnum = "id"
	ListDataEntitiesSortByTimecreated ListDataEntitiesSortByEnum = "timeCreated"
	ListDataEntitiesSortByDisplayname ListDataEntitiesSortByEnum = "displayName"
)

var mappingListDataEntitiesSortByEnum = map[string]ListDataEntitiesSortByEnum{
	"id":          ListDataEntitiesSortById,
	"timeCreated": ListDataEntitiesSortByTimecreated,
	"displayName": ListDataEntitiesSortByDisplayname,
}

var mappingListDataEntitiesSortByEnumLowerCase = map[string]ListDataEntitiesSortByEnum{
	"id":          ListDataEntitiesSortById,
	"timecreated": ListDataEntitiesSortByTimecreated,
	"displayname": ListDataEntitiesSortByDisplayname,
}

// GetListDataEntitiesSortByEnumValues Enumerates the set of values for ListDataEntitiesSortByEnum
func GetListDataEntitiesSortByEnumValues() []ListDataEntitiesSortByEnum {
	values := make([]ListDataEntitiesSortByEnum, 0)
	for _, v := range mappingListDataEntitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataEntitiesSortByEnumStringValues Enumerates the set of values in String for ListDataEntitiesSortByEnum
func GetListDataEntitiesSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDataEntitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataEntitiesSortByEnum(val string) (ListDataEntitiesSortByEnum, bool) {
	enum, ok := mappingListDataEntitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataEntitiesSortOrderEnum Enum with underlying type: string
type ListDataEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListDataEntitiesSortOrderEnum
const (
	ListDataEntitiesSortOrderAsc  ListDataEntitiesSortOrderEnum = "ASC"
	ListDataEntitiesSortOrderDesc ListDataEntitiesSortOrderEnum = "DESC"
)

var mappingListDataEntitiesSortOrderEnum = map[string]ListDataEntitiesSortOrderEnum{
	"ASC":  ListDataEntitiesSortOrderAsc,
	"DESC": ListDataEntitiesSortOrderDesc,
}

var mappingListDataEntitiesSortOrderEnumLowerCase = map[string]ListDataEntitiesSortOrderEnum{
	"asc":  ListDataEntitiesSortOrderAsc,
	"desc": ListDataEntitiesSortOrderDesc,
}

// GetListDataEntitiesSortOrderEnumValues Enumerates the set of values for ListDataEntitiesSortOrderEnum
func GetListDataEntitiesSortOrderEnumValues() []ListDataEntitiesSortOrderEnum {
	values := make([]ListDataEntitiesSortOrderEnum, 0)
	for _, v := range mappingListDataEntitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataEntitiesSortOrderEnumStringValues Enumerates the set of values in String for ListDataEntitiesSortOrderEnum
func GetListDataEntitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataEntitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataEntitiesSortOrderEnum(val string) (ListDataEntitiesSortOrderEnum, bool) {
	enum, ok := mappingListDataEntitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataEntitiesApiModeEnum Enum with underlying type: string
type ListDataEntitiesApiModeEnum string

// Set of constants representing the allowable values for ListDataEntitiesApiModeEnum
const (
	ListDataEntitiesApiModePreview ListDataEntitiesApiModeEnum = "PREVIEW"
	ListDataEntitiesApiModeAll     ListDataEntitiesApiModeEnum = "ALL"
)

var mappingListDataEntitiesApiModeEnum = map[string]ListDataEntitiesApiModeEnum{
	"PREVIEW": ListDataEntitiesApiModePreview,
	"ALL":     ListDataEntitiesApiModeAll,
}

var mappingListDataEntitiesApiModeEnumLowerCase = map[string]ListDataEntitiesApiModeEnum{
	"preview": ListDataEntitiesApiModePreview,
	"all":     ListDataEntitiesApiModeAll,
}

// GetListDataEntitiesApiModeEnumValues Enumerates the set of values for ListDataEntitiesApiModeEnum
func GetListDataEntitiesApiModeEnumValues() []ListDataEntitiesApiModeEnum {
	values := make([]ListDataEntitiesApiModeEnum, 0)
	for _, v := range mappingListDataEntitiesApiModeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataEntitiesApiModeEnumStringValues Enumerates the set of values in String for ListDataEntitiesApiModeEnum
func GetListDataEntitiesApiModeEnumStringValues() []string {
	return []string{
		"PREVIEW",
		"ALL",
	}
}

// GetMappingListDataEntitiesApiModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataEntitiesApiModeEnum(val string) (ListDataEntitiesApiModeEnum, bool) {
	enum, ok := mappingListDataEntitiesApiModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
