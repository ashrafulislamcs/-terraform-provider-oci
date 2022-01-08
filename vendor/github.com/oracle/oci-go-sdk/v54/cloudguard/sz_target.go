// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard API
//
// Use the Cloud Guard API to automate processes that you would otherwise perform through the Cloud Guard Console.
// **Note:** You can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v54/common"
	"strings"
)

// SzTarget The information about Security Zone Target.
type SzTarget struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Resource ID which the target uses to monitor
	TargetResourceId *string `mandatory:"true" json:"targetResourceId"`

	// Total number of recipes attached to target
	RecipeCount *int `mandatory:"true" json:"recipeCount"`

	// Target display name, can be renamed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The target description.
	Description *string `mandatory:"false" json:"description"`

	// List of detector recipes associated with target
	TargetDetectorRecipes []TargetDetectorRecipe `mandatory:"false" json:"targetDetectorRecipes"`

	// List of responder recipes associated with target
	TargetResponderRecipes []TargetResponderRecipe `mandatory:"false" json:"targetResponderRecipes"`

	// List of inherited compartments
	InheritedByCompartments []string `mandatory:"false" json:"inheritedByCompartments"`

	// The date and time the target was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the target was updated. Format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The OCID of the Security Zone to associate this compartment with.
	SecurityZoneId *string `mandatory:"false" json:"securityZoneId"`

	// The name of security zone to associate this compartment with.
	SecurityZoneDisplayName *string `mandatory:"false" json:"securityZoneDisplayName"`

	// Target securityZone recipe list.
	TargetSecurityZoneRecipes []SecurityRecipe `mandatory:"false" json:"targetSecurityZoneRecipes"`

	// The current state of the Target.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m SzTarget) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m SzTarget) GetDisplayName() *string {
	return m.DisplayName
}

//GetCompartmentId returns CompartmentId
func (m SzTarget) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDescription returns Description
func (m SzTarget) GetDescription() *string {
	return m.Description
}

//GetTargetResourceId returns TargetResourceId
func (m SzTarget) GetTargetResourceId() *string {
	return m.TargetResourceId
}

//GetRecipeCount returns RecipeCount
func (m SzTarget) GetRecipeCount() *int {
	return m.RecipeCount
}

//GetTargetDetectorRecipes returns TargetDetectorRecipes
func (m SzTarget) GetTargetDetectorRecipes() []TargetDetectorRecipe {
	return m.TargetDetectorRecipes
}

//GetTargetResponderRecipes returns TargetResponderRecipes
func (m SzTarget) GetTargetResponderRecipes() []TargetResponderRecipe {
	return m.TargetResponderRecipes
}

//GetInheritedByCompartments returns InheritedByCompartments
func (m SzTarget) GetInheritedByCompartments() []string {
	return m.InheritedByCompartments
}

//GetTimeCreated returns TimeCreated
func (m SzTarget) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m SzTarget) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m SzTarget) GetLifecycleState() LifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecyleDetails returns LifecyleDetails
func (m SzTarget) GetLifecyleDetails() *string {
	return m.LifecyleDetails
}

//GetFreeformTags returns FreeformTags
func (m SzTarget) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m SzTarget) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m SzTarget) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m SzTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SzTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SzTarget) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSzTarget SzTarget
	s := struct {
		DiscriminatorParam string `json:"targetResourceType"`
		MarshalTypeSzTarget
	}{
		"SECURITY_ZONE",
		(MarshalTypeSzTarget)(m),
	}

	return json.Marshal(&s)
}
