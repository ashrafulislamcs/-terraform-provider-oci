---
subcategory: "Web Application Acceleration and Security"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_custom_protection_rule"
sidebar_current: "docs-oci-resource-waas-custom_protection_rule"
description: |-
  Provides the Custom Protection Rule resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service
---

# oci_waas_custom_protection_rule
This resource provides the Custom Protection Rule resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service.

Creates a new custom protection rule in the specified compartment.

Custom protection rules allow you to create rules in addition to the rulesets provided by the Web Application Firewall service, including rules from [ModSecurity](https://modsecurity.org/). The syntax for custom rules is based on the ModSecurity syntax. For more information about custom protection rules, see [Custom Protection Rules](https://docs.cloud.oracle.com/iaas/Content/WAF/Tasks/customprotectionrules.htm).

## Example Usage

```hcl
resource "oci_waas_custom_protection_rule" "test_custom_protection_rule" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.custom_protection_rule_display_name
	template = var.custom_protection_rule_template

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.custom_protection_rule_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the custom protection rule.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A description for the Custom Protection rule.
* `display_name` - (Required) (Updatable) A user-friendly name for the custom protection rule.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `template` - (Required) (Updatable) The template text of the custom protection rule. All custom protection rules are expressed in ModSecurity Rule Language.

	Additionally, each rule must include two placeholder variables that are updated by the WAF service upon publication of the rule.

	`id: {{id_1}}` - This field is populated with a unique rule ID generated by the WAF service which identifies a `SecRule`. More than one `SecRule` can be defined in the `template` field of a CreateCustomSecurityRule call. The value of the first `SecRule` must be `id: {{id_1}}` and the `id` field of each subsequent `SecRule` should increase by one, as shown in the example.

	`ctl:ruleEngine={{mode}}` - The action to be taken when the criteria of the `SecRule` are met, either `OFF`, `DETECT` or `BLOCK`. This field is automatically populated with the corresponding value of the `action` field of the `CustomProtectionRuleSetting` schema when the `WafConfig` is updated.

	*Example:* ``` SecRule REQUEST_COOKIES "regex matching SQL injection - part 1/2" \ "phase:2,                                                 \ msg:'Detects chained SQL injection attempts 1/2.',        \ id: {{id_1}},                                             \ ctl:ruleEngine={{mode}},                                  \ deny" SecRule REQUEST_COOKIES "regex matching SQL injection - part 2/2" \ "phase:2,                                                 \ msg:'Detects chained SQL injection attempts 2/2.',        \ id: {{id_2}},                                             \ ctl:ruleEngine={{mode}},                                  \ deny" ```

	 The example contains two `SecRules` each having distinct regex expression to match the `Cookie` header value during the second input analysis phase.

	For more information about custom protection rules, see [Custom Protection Rules](https://docs.cloud.oracle.com/iaas/Content/WAF/tasks/customprotectionrules.htm).

	For more information about ModSecurity syntax, see [Making Rules: The Basic Syntax](https://www.modsecurity.org/CRS/Documentation/making.html).

	For more information about ModSecurity's open source WAF rules, see [Mod Security's OWASP Core Rule Set documentation](https://www.modsecurity.org/CRS/Documentation/index.html).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom protection rule's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the custom protection rule.
* `display_name` - The user-friendly name of the custom protection rule.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom protection rule.
* `mod_security_rule_ids` - The auto-generated ID for the custom protection rule. These IDs are referenced in logs.
* `state` - The current lifecycle state of the custom protection rule.
* `template` - The template text of the custom protection rule. All custom protection rules are expressed in ModSecurity Rule Language.

	Additionally, each rule must include two placeholder variables that are updated by the WAF service upon publication of the rule.

	`id: {{id_1}}` - This field is populated with a unique rule ID generated by the WAF service which identifies a `SecRule`. More than one `SecRule` can be defined in the `template` field of a CreateCustomSecurityRule call. The value of the first `SecRule` must be `id: {{id_1}}` and the `id` field of each subsequent `SecRule` should increase by one, as shown in the example.

	`ctl:ruleEngine={{mode}}` - The action to be taken when the criteria of the `SecRule` are met, either `OFF`, `DETECT` or `BLOCK`. This field is automatically populated with the corresponding value of the `action` field of the `CustomProtectionRuleSetting` schema when the `WafConfig` is updated.

	*Example:* ``` SecRule REQUEST_COOKIES "regex matching SQL injection - part 1/2" \ "phase:2,                                                 \ msg:'Detects chained SQL injection attempts 1/2.',        \ id: {{id_1}},                                             \ ctl:ruleEngine={{mode}},                                  \ deny" SecRule REQUEST_COOKIES "regex matching SQL injection - part 2/2" \ "phase:2,                                                 \ msg:'Detects chained SQL injection attempts 2/2.',        \ id: {{id_2}},                                             \ ctl:ruleEngine={{mode}},                                  \ deny" ```

	 The example contains two `SecRules` each having distinct regex expression to match the `Cookie` header value during the second input analysis phase.

	For more information about custom protection rules, see [Custom Protection Rules](https://docs.cloud.oracle.com/iaas/Content/WAF/tasks/customprotectionrules.htm).

	For more information about ModSecurity syntax, see [Making Rules: The Basic Syntax](https://www.modsecurity.org/CRS/Documentation/making.html).

	For more information about ModSecurity's open source WAF rules, see [Mod Security's OWASP Core Rule Set documentation](https://www.modsecurity.org/CRS/Documentation/index.html).
* `time_created` - The date and time the protection rule was created, expressed in RFC 3339 timestamp format.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Custom Protection Rule
	* `update` - (Defaults to 20 minutes), when updating the Custom Protection Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Custom Protection Rule


## Import

CustomProtectionRules can be imported using the `id`, e.g.

```
$ terraform import oci_waas_custom_protection_rule.test_custom_protection_rule "id"
```

