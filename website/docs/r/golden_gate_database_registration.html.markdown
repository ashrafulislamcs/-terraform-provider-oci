---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_database_registration"
sidebar_current: "docs-oci-resource-golden_gate-database_registration"
description: |-
  Provides the Database Registration resource in Oracle Cloud Infrastructure Golden Gate service
---

# oci_golden_gate_database_registration
This resource provides the Database Registration resource in Oracle Cloud Infrastructure Golden Gate service.

Creates a new DatabaseRegistration.


## Example Usage

```hcl
resource "oci_golden_gate_database_registration" "test_database_registration" {
	#Required
	alias_name = var.database_registration_alias_name
	compartment_id = var.compartment_id
	display_name = var.database_registration_display_name
	fqdn = var.database_registration_fqdn
	password = var.database_registration_password
	username = var.database_registration_username

	#Optional
	connection_string = var.database_registration_connection_string
	database_id = oci_database_database.test_database.id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.database_registration_description
	freeform_tags = {"bar-key"= "value"}
	ip_address = var.database_registration_ip_address
	key_id = oci_kms_key.test_key.id
	secret_compartment_id = oci_identity_compartment.test_compartment.id
	subnet_id = oci_core_subnet.test_subnet.id
	vault_id = oci_kms_vault.test_vault.id
	wallet = var.database_registration_wallet
}
```

## Argument Reference

The following arguments are supported:

* `alias_name` - (Required) (Updatable) Credential store alias. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `connection_string` - (Optional) (Updatable) Connect descriptor or Easy Connect Naming method that Oracle GoldenGate uses to connect to a database. 
* `database_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database being referenced. 
* `defined_tags` - (Optional) (Updatable) Tags defined for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Metadata about this specific object. 
* `display_name` - (Required) (Updatable) An object's Display Name. 
* `fqdn` - (Required) (Updatable) A three-label Fully Qualified Domain Name (FQDN) for a resource. 
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `ip_address` - (Optional) The private IP address in the customer's VCN of the customer's endpoint, typically a database. 
* `key_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer "Master" key being referenced. If provided, this will reference a key which the customer will be required to ensure the policies are established to permit the GoldenGate Service to utilize this key to manage secrets. 
* `password` - (Required) (Updatable) The password Oracle GoldenGate uses to connect the associated RDBMS.  It must conform to the specific security requirements implemented by the database including length, case sensitivity, and so on. 
* `secret_compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the the GGS Secret will be created. If provided, this will reference a key which the customer will be required to ensure the policies are established to permit the GoldenGate Service to utilize this Compartment in which to create a Secret. 
* `subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet being referenced. 
* `username` - (Required) (Updatable) The username Oracle GoldenGate uses to connect the associated RDBMS.  This username must already exist and be available for use by the database.  It must conform to the security requirements implemented by the database including length, case sensitivity, and so on. 
* `vault_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer vault being referenced. If provided, this will reference a vault which the customer will be required to ensure the policies are established to permit the GoldenGate Service to manage secrets contained within this vault. 
* `wallet` - (Optional) (Updatable) The wallet contents Oracle GoldenGate uses to make connections to a database.  This attribute is expected to be base64 encoded. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `alias_name` - Credential store alias. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced. 
* `connection_string` - Connect descriptor or Easy Connect Naming method that Oracle GoldenGate uses to connect to a database. 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database being referenced. 
* `defined_tags` - Tags defined for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Metadata about this specific object. 
* `display_name` - An object's Display Name. 
* `fqdn` - A three-label Fully Qualified Domain Name (FQDN) for a resource. 
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the databaseRegistration being referenced. 
* `ip_address` - The private IP address in the customer's VCN of the customer's endpoint, typically a database. 
* `key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer "Master" key being referenced. If provided, this will reference a key which the customer will be required to ensure the policies are established to permit the GoldenGate Service to utilize this key to manage secrets. 
* `lifecycle_details` - Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state. 
* `rce_private_ip` - A Private Endpoint IP Address created in the customer's subnet.  A customer database can expect network traffic initiated by GGS from this IP address and send network traffic to this IP address, typically in response to requests from GGS (OGG).  The customer may utilize this IP address in Security Lists or Network Security Groups (NSG) as needed. 
* `secret_compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the the GGS Secret will be created. If provided, this will reference a key which the customer will be required to ensure the policies are established to permit the GoldenGate Service to utilize this Compartment in which to create a Secret. 
* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer GGS Secret being referenced. If provided, this will reference a key which the customer will be required to ensure the policies are established to permit the GoldenGate Service to utilize this Secret 
* `state` - Possible lifecycle states. 
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet being referenced. 
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time the resource was created. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `time_updated` - The time the resource was last updated. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`. 
* `username` - The username Oracle GoldenGate uses to connect the associated RDBMS.  This username must already exist and be available for use by the database.  It must conform to the security requirements implemented by the database including length, case sensitivity, and so on. 
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer vault being referenced. If provided, this will reference a vault which the customer will be required to ensure the policies are established to permit the GoldenGate Service to manage secrets contained within this vault. 

## Import

DatabaseRegistrations can be imported using the `id`, e.g.

```
$ terraform import oci_golden_gate_database_registration.test_database_registration "id"
```
