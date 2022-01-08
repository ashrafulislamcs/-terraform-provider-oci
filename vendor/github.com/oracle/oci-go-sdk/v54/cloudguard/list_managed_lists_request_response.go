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

// ListManagedListsRequest wrapper for the ListManagedLists operation
type ListManagedListsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Default is false.
	// When set to true, the list of all Oracle Managed Resources
	// Metadata supported by Cloud Guard are returned.
	ResourceMetadataOnly *bool `mandatory:"false" contributesTo:"query" name:"resourceMetadataOnly"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListManagedListsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The type of the ManagedList.
	ListType ListManagedListsListTypeEnum `mandatory:"false" contributesTo:"query" name:"listType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListManagedListsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagedListsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListManagedListsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedListsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedListsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedListsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedListsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedListsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingListManagedListsLifecycleStateEnum[string(request.LifecycleState)]; !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListManagedListsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := mappingListManagedListsListTypeEnum[string(request.ListType)]; !ok && request.ListType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListType: %s. Supported values are: %s.", request.ListType, strings.Join(GetListManagedListsListTypeEnumStringValues(), ",")))
	}
	if _, ok := mappingListManagedListsAccessLevelEnum[string(request.AccessLevel)]; !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListManagedListsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := mappingListManagedListsSortOrderEnum[string(request.SortOrder)]; !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedListsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := mappingListManagedListsSortByEnum[string(request.SortBy)]; !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedListsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedListsResponse wrapper for the ListManagedLists operation
type ListManagedListsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedListCollection instances
	ManagedListCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedListsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedListsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedListsLifecycleStateEnum Enum with underlying type: string
type ListManagedListsLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagedListsLifecycleStateEnum
const (
	ListManagedListsLifecycleStateCreating ListManagedListsLifecycleStateEnum = "CREATING"
	ListManagedListsLifecycleStateUpdating ListManagedListsLifecycleStateEnum = "UPDATING"
	ListManagedListsLifecycleStateActive   ListManagedListsLifecycleStateEnum = "ACTIVE"
	ListManagedListsLifecycleStateInactive ListManagedListsLifecycleStateEnum = "INACTIVE"
	ListManagedListsLifecycleStateDeleting ListManagedListsLifecycleStateEnum = "DELETING"
	ListManagedListsLifecycleStateDeleted  ListManagedListsLifecycleStateEnum = "DELETED"
	ListManagedListsLifecycleStateFailed   ListManagedListsLifecycleStateEnum = "FAILED"
)

var mappingListManagedListsLifecycleStateEnum = map[string]ListManagedListsLifecycleStateEnum{
	"CREATING": ListManagedListsLifecycleStateCreating,
	"UPDATING": ListManagedListsLifecycleStateUpdating,
	"ACTIVE":   ListManagedListsLifecycleStateActive,
	"INACTIVE": ListManagedListsLifecycleStateInactive,
	"DELETING": ListManagedListsLifecycleStateDeleting,
	"DELETED":  ListManagedListsLifecycleStateDeleted,
	"FAILED":   ListManagedListsLifecycleStateFailed,
}

// GetListManagedListsLifecycleStateEnumValues Enumerates the set of values for ListManagedListsLifecycleStateEnum
func GetListManagedListsLifecycleStateEnumValues() []ListManagedListsLifecycleStateEnum {
	values := make([]ListManagedListsLifecycleStateEnum, 0)
	for _, v := range mappingListManagedListsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedListsLifecycleStateEnumStringValues Enumerates the set of values in String for ListManagedListsLifecycleStateEnum
func GetListManagedListsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// ListManagedListsListTypeEnum Enum with underlying type: string
type ListManagedListsListTypeEnum string

// Set of constants representing the allowable values for ListManagedListsListTypeEnum
const (
	ListManagedListsListTypeCidrBlock    ListManagedListsListTypeEnum = "CIDR_BLOCK"
	ListManagedListsListTypeUsers        ListManagedListsListTypeEnum = "USERS"
	ListManagedListsListTypeGroups       ListManagedListsListTypeEnum = "GROUPS"
	ListManagedListsListTypeIpv4address  ListManagedListsListTypeEnum = "IPV4ADDRESS"
	ListManagedListsListTypeIpv6address  ListManagedListsListTypeEnum = "IPV6ADDRESS"
	ListManagedListsListTypeResourceOcid ListManagedListsListTypeEnum = "RESOURCE_OCID"
	ListManagedListsListTypeRegion       ListManagedListsListTypeEnum = "REGION"
	ListManagedListsListTypeCountry      ListManagedListsListTypeEnum = "COUNTRY"
	ListManagedListsListTypeState        ListManagedListsListTypeEnum = "STATE"
	ListManagedListsListTypeCity         ListManagedListsListTypeEnum = "CITY"
	ListManagedListsListTypeTags         ListManagedListsListTypeEnum = "TAGS"
	ListManagedListsListTypeGeneric      ListManagedListsListTypeEnum = "GENERIC"
)

var mappingListManagedListsListTypeEnum = map[string]ListManagedListsListTypeEnum{
	"CIDR_BLOCK":    ListManagedListsListTypeCidrBlock,
	"USERS":         ListManagedListsListTypeUsers,
	"GROUPS":        ListManagedListsListTypeGroups,
	"IPV4ADDRESS":   ListManagedListsListTypeIpv4address,
	"IPV6ADDRESS":   ListManagedListsListTypeIpv6address,
	"RESOURCE_OCID": ListManagedListsListTypeResourceOcid,
	"REGION":        ListManagedListsListTypeRegion,
	"COUNTRY":       ListManagedListsListTypeCountry,
	"STATE":         ListManagedListsListTypeState,
	"CITY":          ListManagedListsListTypeCity,
	"TAGS":          ListManagedListsListTypeTags,
	"GENERIC":       ListManagedListsListTypeGeneric,
}

// GetListManagedListsListTypeEnumValues Enumerates the set of values for ListManagedListsListTypeEnum
func GetListManagedListsListTypeEnumValues() []ListManagedListsListTypeEnum {
	values := make([]ListManagedListsListTypeEnum, 0)
	for _, v := range mappingListManagedListsListTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedListsListTypeEnumStringValues Enumerates the set of values in String for ListManagedListsListTypeEnum
func GetListManagedListsListTypeEnumStringValues() []string {
	return []string{
		"CIDR_BLOCK",
		"USERS",
		"GROUPS",
		"IPV4ADDRESS",
		"IPV6ADDRESS",
		"RESOURCE_OCID",
		"REGION",
		"COUNTRY",
		"STATE",
		"CITY",
		"TAGS",
		"GENERIC",
	}
}

// ListManagedListsAccessLevelEnum Enum with underlying type: string
type ListManagedListsAccessLevelEnum string

// Set of constants representing the allowable values for ListManagedListsAccessLevelEnum
const (
	ListManagedListsAccessLevelRestricted ListManagedListsAccessLevelEnum = "RESTRICTED"
	ListManagedListsAccessLevelAccessible ListManagedListsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListManagedListsAccessLevelEnum = map[string]ListManagedListsAccessLevelEnum{
	"RESTRICTED": ListManagedListsAccessLevelRestricted,
	"ACCESSIBLE": ListManagedListsAccessLevelAccessible,
}

// GetListManagedListsAccessLevelEnumValues Enumerates the set of values for ListManagedListsAccessLevelEnum
func GetListManagedListsAccessLevelEnumValues() []ListManagedListsAccessLevelEnum {
	values := make([]ListManagedListsAccessLevelEnum, 0)
	for _, v := range mappingListManagedListsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedListsAccessLevelEnumStringValues Enumerates the set of values in String for ListManagedListsAccessLevelEnum
func GetListManagedListsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// ListManagedListsSortOrderEnum Enum with underlying type: string
type ListManagedListsSortOrderEnum string

// Set of constants representing the allowable values for ListManagedListsSortOrderEnum
const (
	ListManagedListsSortOrderAsc  ListManagedListsSortOrderEnum = "ASC"
	ListManagedListsSortOrderDesc ListManagedListsSortOrderEnum = "DESC"
)

var mappingListManagedListsSortOrderEnum = map[string]ListManagedListsSortOrderEnum{
	"ASC":  ListManagedListsSortOrderAsc,
	"DESC": ListManagedListsSortOrderDesc,
}

// GetListManagedListsSortOrderEnumValues Enumerates the set of values for ListManagedListsSortOrderEnum
func GetListManagedListsSortOrderEnumValues() []ListManagedListsSortOrderEnum {
	values := make([]ListManagedListsSortOrderEnum, 0)
	for _, v := range mappingListManagedListsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedListsSortOrderEnumStringValues Enumerates the set of values in String for ListManagedListsSortOrderEnum
func GetListManagedListsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// ListManagedListsSortByEnum Enum with underlying type: string
type ListManagedListsSortByEnum string

// Set of constants representing the allowable values for ListManagedListsSortByEnum
const (
	ListManagedListsSortByTimecreated ListManagedListsSortByEnum = "timeCreated"
	ListManagedListsSortByDisplayname ListManagedListsSortByEnum = "displayName"
)

var mappingListManagedListsSortByEnum = map[string]ListManagedListsSortByEnum{
	"timeCreated": ListManagedListsSortByTimecreated,
	"displayName": ListManagedListsSortByDisplayname,
}

// GetListManagedListsSortByEnumValues Enumerates the set of values for ListManagedListsSortByEnum
func GetListManagedListsSortByEnumValues() []ListManagedListsSortByEnum {
	values := make([]ListManagedListsSortByEnum, 0)
	for _, v := range mappingListManagedListsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedListsSortByEnumStringValues Enumerates the set of values in String for ListManagedListsSortByEnum
func GetListManagedListsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}
