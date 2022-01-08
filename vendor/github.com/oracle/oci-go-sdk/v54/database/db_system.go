// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v54/common"
	"strings"
)

// DbSystem The representation of DbSystem
type DbSystem struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the DB system. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the availability domain that the DB system is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the DB system is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and backup subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The shape of the DB system. The shape determines resources to allocate to the DB system.
	// - For virtual machine shapes, the number of CPU cores and memory
	// - For bare metal and Exadata shapes, the number of CPU cores, storage, and memory
	Shape *string `mandatory:"true" json:"shape"`

	// The public key portion of one or more key pairs used for SSH access to the DB system.
	SshPublicKeys []string `mandatory:"true" json:"sshPublicKeys"`

	// The hostname for the DB system.
	Hostname *string `mandatory:"true" json:"hostname"`

	// The domain name for the DB system.
	Domain *string `mandatory:"true" json:"domain"`

	// The number of CPU cores enabled on the DB system.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The Oracle Database edition that applies to all the databases on the DB system.
	DatabaseEdition DbSystemDatabaseEditionEnum `mandatory:"true" json:"databaseEdition"`

	// The current state of the DB system.
	LifecycleState DbSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// List of the Fault Domains in which this DB system is provisioned.
	FaultDomains []string `mandatory:"false" json:"faultDomains"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup network subnet the DB system is associated with. Applicable only to Exadata DB systems.
	// **Subnet Restriction:** See the subnet restrictions information for **subnetId**.
	BackupSubnetId *string `mandatory:"false" json:"backupSubnetId"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that the backup network of this DB system belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm). Applicable only to Exadata systems.
	BackupNetworkNsgIds []string `mandatory:"false" json:"backupNetworkNsgIds"`

	DbSystemOptions *DbSystemOptions `mandatory:"false" json:"dbSystemOptions"`

	// The time zone of the DB system. For details, see DB System Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The Oracle Database version of the DB system.
	Version *string `mandatory:"false" json:"version"`

	// The cluster name for Exadata and 2-node RAC virtual machine DB systems. The cluster name must begin with an alphabetic character, and may contain hyphens (-). Underscores (_) are not permitted. The cluster name can be no longer than 11 characters and is not case sensitive.
	ClusterName *string `mandatory:"false" json:"clusterName"`

	// The percentage assigned to DATA storage (user data and database files).
	// The remaining percentage is assigned to RECO storage (database redo logs, archive logs, and recovery manager backups). Accepted values are 40 and 80. The default is 80 percent assigned to DATA storage. Not applicable for virtual machine DB systems.
	DataStoragePercentage *int `mandatory:"false" json:"dataStoragePercentage"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation starts.
	LastPatchHistoryEntryId *string `mandatory:"false" json:"lastPatchHistoryEntryId"`

	// The port number configured for the listener on the DB system.
	ListenerPort *int `mandatory:"false" json:"listenerPort"`

	// The date and time the DB system was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The type of redundancy configured for the DB system.
	// NORMAL is 2-way redundancy.
	// HIGH is 3-way redundancy.
	DiskRedundancy DbSystemDiskRedundancyEnum `mandatory:"false" json:"diskRedundancy,omitempty"`

	// True, if Sparse Diskgroup is configured for Exadata dbsystem, False, if Sparse diskgroup was not configured.
	SparseDiskgroup *bool `mandatory:"false" json:"sparseDiskgroup"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Single Client Access Name (SCAN) IP addresses associated with the DB system.
	// SCAN IP addresses are typically used for load balancing and are not assigned to any interface.
	// Oracle Clusterware directs the requests to the appropriate nodes in the cluster.
	// **Note:** For a single-node DB system, this list is empty.
	ScanIpIds []string `mandatory:"false" json:"scanIpIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the virtual IP (VIP) addresses associated with the DB system.
	// The Cluster Ready Services (CRS) creates and maintains one VIP address for each node in the DB system to
	// enable failover. If one node fails, the VIP is reassigned to another active node in the cluster.
	// **Note:** For a single-node DB system, this list is empty.
	VipIds []string `mandatory:"false" json:"vipIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DNS record for the SCAN IP addresses that are associated with the DB system.
	ScanDnsRecordId *string `mandatory:"false" json:"scanDnsRecordId"`

	// The FQDN of the DNS record for the SCAN IP addresses that are associated with the DB system.
	ScanDnsName *string `mandatory:"false" json:"scanDnsName"`

	// The OCID of the zone the DB system is associated with.
	ZoneId *string `mandatory:"false" json:"zoneId"`

	// The data storage size, in gigabytes, that is currently available to the DB system. Applies only for virtual machine DB systems.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The RECO/REDO storage size, in gigabytes, that is currently allocated to the DB system. Applies only for virtual machine DB systems.
	RecoStorageSizeInGB *int `mandatory:"false" json:"recoStorageSizeInGB"`

	// The number of nodes in the DB system. For RAC DB systems, the value is greater than 1.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The Oracle license model that applies to all the databases on the DB system. The default is LICENSE_INCLUDED.
	LicenseModel DbSystemLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last maintenance run.
	LastMaintenanceRunId *string `mandatory:"false" json:"lastMaintenanceRunId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the next maintenance run.
	NextMaintenanceRunId *string `mandatory:"false" json:"nextMaintenanceRunId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the DB system.
	SourceDbSystemId *string `mandatory:"false" json:"sourceDbSystemId"`

	// The point in time for a cloned database system when the data disks were cloned from the source database system, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	PointInTimeDataDiskCloneTimestamp *common.SDKTime `mandatory:"false" json:"pointInTimeDataDiskCloneTimestamp"`

	IormConfigCache *ExadataIormConfig `mandatory:"false" json:"iormConfigCache"`
}

func (m DbSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingDbSystemDatabaseEditionEnum[string(m.DatabaseEdition)]; !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetDbSystemDatabaseEditionEnumStringValues(), ",")))
	}
	if _, ok := mappingDbSystemLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbSystemLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := mappingDbSystemDiskRedundancyEnum[string(m.DiskRedundancy)]; !ok && m.DiskRedundancy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiskRedundancy: %s. Supported values are: %s.", m.DiskRedundancy, strings.Join(GetDbSystemDiskRedundancyEnumStringValues(), ",")))
	}
	if _, ok := mappingDbSystemLicenseModelEnum[string(m.LicenseModel)]; !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetDbSystemLicenseModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbSystemDatabaseEditionEnum Enum with underlying type: string
type DbSystemDatabaseEditionEnum string

// Set of constants representing the allowable values for DbSystemDatabaseEditionEnum
const (
	DbSystemDatabaseEditionStandardEdition                     DbSystemDatabaseEditionEnum = "STANDARD_EDITION"
	DbSystemDatabaseEditionEnterpriseEdition                   DbSystemDatabaseEditionEnum = "ENTERPRISE_EDITION"
	DbSystemDatabaseEditionEnterpriseEditionHighPerformance    DbSystemDatabaseEditionEnum = "ENTERPRISE_EDITION_HIGH_PERFORMANCE"
	DbSystemDatabaseEditionEnterpriseEditionExtremePerformance DbSystemDatabaseEditionEnum = "ENTERPRISE_EDITION_EXTREME_PERFORMANCE"
)

var mappingDbSystemDatabaseEditionEnum = map[string]DbSystemDatabaseEditionEnum{
	"STANDARD_EDITION":                       DbSystemDatabaseEditionStandardEdition,
	"ENTERPRISE_EDITION":                     DbSystemDatabaseEditionEnterpriseEdition,
	"ENTERPRISE_EDITION_HIGH_PERFORMANCE":    DbSystemDatabaseEditionEnterpriseEditionHighPerformance,
	"ENTERPRISE_EDITION_EXTREME_PERFORMANCE": DbSystemDatabaseEditionEnterpriseEditionExtremePerformance,
}

// GetDbSystemDatabaseEditionEnumValues Enumerates the set of values for DbSystemDatabaseEditionEnum
func GetDbSystemDatabaseEditionEnumValues() []DbSystemDatabaseEditionEnum {
	values := make([]DbSystemDatabaseEditionEnum, 0)
	for _, v := range mappingDbSystemDatabaseEditionEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemDatabaseEditionEnumStringValues Enumerates the set of values in String for DbSystemDatabaseEditionEnum
func GetDbSystemDatabaseEditionEnumStringValues() []string {
	return []string{
		"STANDARD_EDITION",
		"ENTERPRISE_EDITION",
		"ENTERPRISE_EDITION_HIGH_PERFORMANCE",
		"ENTERPRISE_EDITION_EXTREME_PERFORMANCE",
	}
}

// DbSystemLifecycleStateEnum Enum with underlying type: string
type DbSystemLifecycleStateEnum string

// Set of constants representing the allowable values for DbSystemLifecycleStateEnum
const (
	DbSystemLifecycleStateProvisioning          DbSystemLifecycleStateEnum = "PROVISIONING"
	DbSystemLifecycleStateAvailable             DbSystemLifecycleStateEnum = "AVAILABLE"
	DbSystemLifecycleStateUpdating              DbSystemLifecycleStateEnum = "UPDATING"
	DbSystemLifecycleStateTerminating           DbSystemLifecycleStateEnum = "TERMINATING"
	DbSystemLifecycleStateTerminated            DbSystemLifecycleStateEnum = "TERMINATED"
	DbSystemLifecycleStateFailed                DbSystemLifecycleStateEnum = "FAILED"
	DbSystemLifecycleStateMigrated              DbSystemLifecycleStateEnum = "MIGRATED"
	DbSystemLifecycleStateMaintenanceInProgress DbSystemLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	DbSystemLifecycleStateNeedsAttention        DbSystemLifecycleStateEnum = "NEEDS_ATTENTION"
	DbSystemLifecycleStateUpgrading             DbSystemLifecycleStateEnum = "UPGRADING"
)

var mappingDbSystemLifecycleStateEnum = map[string]DbSystemLifecycleStateEnum{
	"PROVISIONING":            DbSystemLifecycleStateProvisioning,
	"AVAILABLE":               DbSystemLifecycleStateAvailable,
	"UPDATING":                DbSystemLifecycleStateUpdating,
	"TERMINATING":             DbSystemLifecycleStateTerminating,
	"TERMINATED":              DbSystemLifecycleStateTerminated,
	"FAILED":                  DbSystemLifecycleStateFailed,
	"MIGRATED":                DbSystemLifecycleStateMigrated,
	"MAINTENANCE_IN_PROGRESS": DbSystemLifecycleStateMaintenanceInProgress,
	"NEEDS_ATTENTION":         DbSystemLifecycleStateNeedsAttention,
	"UPGRADING":               DbSystemLifecycleStateUpgrading,
}

// GetDbSystemLifecycleStateEnumValues Enumerates the set of values for DbSystemLifecycleStateEnum
func GetDbSystemLifecycleStateEnumValues() []DbSystemLifecycleStateEnum {
	values := make([]DbSystemLifecycleStateEnum, 0)
	for _, v := range mappingDbSystemLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemLifecycleStateEnumStringValues Enumerates the set of values in String for DbSystemLifecycleStateEnum
func GetDbSystemLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"MIGRATED",
		"MAINTENANCE_IN_PROGRESS",
		"NEEDS_ATTENTION",
		"UPGRADING",
	}
}

// DbSystemDiskRedundancyEnum Enum with underlying type: string
type DbSystemDiskRedundancyEnum string

// Set of constants representing the allowable values for DbSystemDiskRedundancyEnum
const (
	DbSystemDiskRedundancyHigh   DbSystemDiskRedundancyEnum = "HIGH"
	DbSystemDiskRedundancyNormal DbSystemDiskRedundancyEnum = "NORMAL"
)

var mappingDbSystemDiskRedundancyEnum = map[string]DbSystemDiskRedundancyEnum{
	"HIGH":   DbSystemDiskRedundancyHigh,
	"NORMAL": DbSystemDiskRedundancyNormal,
}

// GetDbSystemDiskRedundancyEnumValues Enumerates the set of values for DbSystemDiskRedundancyEnum
func GetDbSystemDiskRedundancyEnumValues() []DbSystemDiskRedundancyEnum {
	values := make([]DbSystemDiskRedundancyEnum, 0)
	for _, v := range mappingDbSystemDiskRedundancyEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemDiskRedundancyEnumStringValues Enumerates the set of values in String for DbSystemDiskRedundancyEnum
func GetDbSystemDiskRedundancyEnumStringValues() []string {
	return []string{
		"HIGH",
		"NORMAL",
	}
}

// DbSystemLicenseModelEnum Enum with underlying type: string
type DbSystemLicenseModelEnum string

// Set of constants representing the allowable values for DbSystemLicenseModelEnum
const (
	DbSystemLicenseModelLicenseIncluded     DbSystemLicenseModelEnum = "LICENSE_INCLUDED"
	DbSystemLicenseModelBringYourOwnLicense DbSystemLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingDbSystemLicenseModelEnum = map[string]DbSystemLicenseModelEnum{
	"LICENSE_INCLUDED":       DbSystemLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": DbSystemLicenseModelBringYourOwnLicense,
}

// GetDbSystemLicenseModelEnumValues Enumerates the set of values for DbSystemLicenseModelEnum
func GetDbSystemLicenseModelEnumValues() []DbSystemLicenseModelEnum {
	values := make([]DbSystemLicenseModelEnum, 0)
	for _, v := range mappingDbSystemLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemLicenseModelEnumStringValues Enumerates the set of values in String for DbSystemLicenseModelEnum
func GetDbSystemLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}
