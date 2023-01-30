---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_snapshots"
sidebar_current: "docs-oci-datasource-file_storage-snapshots"
description: |-
  Provides the list of Snapshots in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_snapshots
This data source provides the list of Snapshots in Oracle Cloud Infrastructure File Storage service.

Lists snapshots of the specified file system.


## Example Usage

```hcl
data "oci_file_storage_snapshots" "test_snapshots" {
	#Required
	file_system_id = oci_file_storage_file_system.test_file_system.id

	#Optional
	id = var.snapshot_id
	state = var.snapshot_state
}
```

## Argument Reference

The following arguments are supported:

* `file_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system.
* `id` - (Optional) Filter results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `snapshots` - The list of snapshots.

### Snapshot Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `file_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system from which the snapshot was created. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the snapshot.
* `is_clone_source` - Specifies whether the snapshot has been cloned. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm). 
* `lifecycle_details` - Additional information about the current `lifecycleState`.
* `name` - Name of the snapshot. This value is immutable.

	Avoid entering confidential information.

	Example: `Sunday` 
* `provenance_id` - An [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) identifying the parent from which this snapshot was cloned. If this snapshot was not cloned, then the `provenanceId` is the same as the snapshot `id` value. If this snapshot was cloned, then the `provenanceId` value is the parent's `provenanceId`. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm). 
* `snapshot_time` - The date and time the snapshot was taken, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format. This value might be the same or different from `timeCreated` depending on the following factors:
	* If the snapshot is created in the original file system directory.
	* If the snapshot is cloned from a file system.
	* If the snapshot is replicated from a file system.

	Example: `2020-08-25T21:10:29.600Z` 
* `snapshot_type` - Specifies the generation type of the snapshot. 
* `state` - The current state of the snapshot.
* `time_created` - The date and time the snapshot was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

