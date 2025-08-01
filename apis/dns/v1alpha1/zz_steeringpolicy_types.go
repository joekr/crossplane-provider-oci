/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AnswerDataInitParameters struct {

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) An expression that is used to select a set of answers that match a condition. For example, answers with matching pool properties.
	AnswerCondition *string `json:"answerCondition,omitempty" tf:"answer_condition,omitempty"`

	// (Applicable when rule_type=FILTER) Keeps the answer only if the value is true.
	ShouldKeep *bool `json:"shouldKeep,omitempty" tf:"should_keep,omitempty"`

	// The rank assigned to the set of answers that match the expression in answerCondition. Answers with the lowest values move to the beginning of the list without changing the relative order of those with the same value. Answers can be given a value between 0 and 255.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type AnswerDataObservation struct {

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) An expression that is used to select a set of answers that match a condition. For example, answers with matching pool properties.
	AnswerCondition *string `json:"answerCondition,omitempty" tf:"answer_condition,omitempty"`

	// (Applicable when rule_type=FILTER) Keeps the answer only if the value is true.
	ShouldKeep *bool `json:"shouldKeep,omitempty" tf:"should_keep,omitempty"`

	// The rank assigned to the set of answers that match the expression in answerCondition. Answers with the lowest values move to the beginning of the list without changing the relative order of those with the same value. Answers can be given a value between 0 and 255.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type AnswerDataParameters struct {

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) An expression that is used to select a set of answers that match a condition. For example, answers with matching pool properties.
	// +kubebuilder:validation:Optional
	AnswerCondition *string `json:"answerCondition,omitempty" tf:"answer_condition,omitempty"`

	// (Applicable when rule_type=FILTER) Keeps the answer only if the value is true.
	// +kubebuilder:validation:Optional
	ShouldKeep *bool `json:"shouldKeep,omitempty" tf:"should_keep,omitempty"`

	// The rank assigned to the set of answers that match the expression in answerCondition. Answers with the lowest values move to the beginning of the list without changing the relative order of those with the same value. Answers can be given a value between 0 and 255.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type AnswersInitParameters struct {

	// Set this property to true to indicate that the answer is administratively disabled, such as when the corresponding server is down for maintenance. An answer's isDisabled property can be referenced in answerCondition properties in rules using answer.isDisabled.
	IsDisabled *bool `json:"isDisabled,omitempty" tf:"is_disabled,omitempty"`

	// A user-friendly name for the answer, unique within the steering policy. An answer's name property can be referenced in answerCondition properties of rules using answer.name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The freeform name of a group of one or more records in which this record is included, such as "LAX data center". An answer's pool property can be referenced in answerCondition properties of rules using answer.pool.
	Pool *string `json:"pool,omitempty" tf:"pool,omitempty"`

	// The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see Supported DNS Resource Record Types.
	Rdata *string `json:"rdata,omitempty" tf:"rdata,omitempty"`

	// The type of DNS record, such as A or CNAME. Only A, AAAA, and CNAME are supported. For more information, see Supported DNS Resource Record Types.
	Rtype *string `json:"rtype,omitempty" tf:"rtype,omitempty"`
}

type AnswersObservation struct {

	// Set this property to true to indicate that the answer is administratively disabled, such as when the corresponding server is down for maintenance. An answer's isDisabled property can be referenced in answerCondition properties in rules using answer.isDisabled.
	IsDisabled *bool `json:"isDisabled,omitempty" tf:"is_disabled,omitempty"`

	// A user-friendly name for the answer, unique within the steering policy. An answer's name property can be referenced in answerCondition properties of rules using answer.name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The freeform name of a group of one or more records in which this record is included, such as "LAX data center". An answer's pool property can be referenced in answerCondition properties of rules using answer.pool.
	Pool *string `json:"pool,omitempty" tf:"pool,omitempty"`

	// The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see Supported DNS Resource Record Types.
	Rdata *string `json:"rdata,omitempty" tf:"rdata,omitempty"`

	// The type of DNS record, such as A or CNAME. Only A, AAAA, and CNAME are supported. For more information, see Supported DNS Resource Record Types.
	Rtype *string `json:"rtype,omitempty" tf:"rtype,omitempty"`
}

type AnswersParameters struct {

	// Set this property to true to indicate that the answer is administratively disabled, such as when the corresponding server is down for maintenance. An answer's isDisabled property can be referenced in answerCondition properties in rules using answer.isDisabled.
	// +kubebuilder:validation:Optional
	IsDisabled *bool `json:"isDisabled,omitempty" tf:"is_disabled,omitempty"`

	// A user-friendly name for the answer, unique within the steering policy. An answer's name property can be referenced in answerCondition properties of rules using answer.name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The freeform name of a group of one or more records in which this record is included, such as "LAX data center". An answer's pool property can be referenced in answerCondition properties of rules using answer.pool.
	// +kubebuilder:validation:Optional
	Pool *string `json:"pool,omitempty" tf:"pool,omitempty"`

	// The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see Supported DNS Resource Record Types.
	// +kubebuilder:validation:Optional
	Rdata *string `json:"rdata" tf:"rdata,omitempty"`

	// The type of DNS record, such as A or CNAME. Only A, AAAA, and CNAME are supported. For more information, see Supported DNS Resource Record Types.
	// +kubebuilder:validation:Optional
	Rtype *string `json:"rtype" tf:"rtype,omitempty"`
}

type CasesInitParameters struct {

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) An array of SteeringPolicyPriorityAnswerData objects.
	AnswerData []AnswerDataInitParameters `json:"answerData,omitempty" tf:"answer_data,omitempty"`

	// (Applicable when rule_type=FILTER | HEALTH | LIMIT | PRIORITY | WEIGHTED) An expression that uses conditions at the time of a DNS query to indicate whether a case matches. Conditions may include the geographical location, IP subnet, or ASN the DNS query originated. Example: If you have an office that uses the subnet 192.0.2.0/24 you could use a caseCondition expression query.client.subnet in ('192.0.2.0/24') to define a case that matches queries from that office.
	CaseCondition *string `json:"caseCondition,omitempty" tf:"case_condition,omitempty"`

	// The number of answers allowed to remain after the limit rule has been processed, keeping only the first of the remaining answers in the list. Example: If the count property is set to 2 and four answers remain before the limit rule is processed, only the first two answers in the list will remain after the limit rule has been processed.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type CasesObservation struct {

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) An array of SteeringPolicyPriorityAnswerData objects.
	AnswerData []AnswerDataObservation `json:"answerData,omitempty" tf:"answer_data,omitempty"`

	// (Applicable when rule_type=FILTER | HEALTH | LIMIT | PRIORITY | WEIGHTED) An expression that uses conditions at the time of a DNS query to indicate whether a case matches. Conditions may include the geographical location, IP subnet, or ASN the DNS query originated. Example: If you have an office that uses the subnet 192.0.2.0/24 you could use a caseCondition expression query.client.subnet in ('192.0.2.0/24') to define a case that matches queries from that office.
	CaseCondition *string `json:"caseCondition,omitempty" tf:"case_condition,omitempty"`

	// The number of answers allowed to remain after the limit rule has been processed, keeping only the first of the remaining answers in the list. Example: If the count property is set to 2 and four answers remain before the limit rule is processed, only the first two answers in the list will remain after the limit rule has been processed.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type CasesParameters struct {

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) An array of SteeringPolicyPriorityAnswerData objects.
	// +kubebuilder:validation:Optional
	AnswerData []AnswerDataParameters `json:"answerData,omitempty" tf:"answer_data,omitempty"`

	// (Applicable when rule_type=FILTER | HEALTH | LIMIT | PRIORITY | WEIGHTED) An expression that uses conditions at the time of a DNS query to indicate whether a case matches. Conditions may include the geographical location, IP subnet, or ASN the DNS query originated. Example: If you have an office that uses the subnet 192.0.2.0/24 you could use a caseCondition expression query.client.subnet in ('192.0.2.0/24') to define a case that matches queries from that office.
	// +kubebuilder:validation:Optional
	CaseCondition *string `json:"caseCondition,omitempty" tf:"case_condition,omitempty"`

	// The number of answers allowed to remain after the limit rule has been processed, keeping only the first of the remaining answers in the list. Example: If the count property is set to 2 and four answers remain before the limit rule is processed, only the first two answers in the list will remain after the limit rule has been processed.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type DefaultAnswerDataInitParameters struct {

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) An expression that is used to select a set of answers that match a condition. For example, answers with matching pool properties.
	AnswerCondition *string `json:"answerCondition,omitempty" tf:"answer_condition,omitempty"`

	// (Applicable when rule_type=FILTER) Keeps the answer only if the value is true.
	ShouldKeep *bool `json:"shouldKeep,omitempty" tf:"should_keep,omitempty"`

	// The rank assigned to the set of answers that match the expression in answerCondition. Answers with the lowest values move to the beginning of the list without changing the relative order of those with the same value. Answers can be given a value between 0 and 255.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type DefaultAnswerDataObservation struct {

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) An expression that is used to select a set of answers that match a condition. For example, answers with matching pool properties.
	AnswerCondition *string `json:"answerCondition,omitempty" tf:"answer_condition,omitempty"`

	// (Applicable when rule_type=FILTER) Keeps the answer only if the value is true.
	ShouldKeep *bool `json:"shouldKeep,omitempty" tf:"should_keep,omitempty"`

	// The rank assigned to the set of answers that match the expression in answerCondition. Answers with the lowest values move to the beginning of the list without changing the relative order of those with the same value. Answers can be given a value between 0 and 255.
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type DefaultAnswerDataParameters struct {

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) An expression that is used to select a set of answers that match a condition. For example, answers with matching pool properties.
	// +kubebuilder:validation:Optional
	AnswerCondition *string `json:"answerCondition,omitempty" tf:"answer_condition,omitempty"`

	// (Applicable when rule_type=FILTER) Keeps the answer only if the value is true.
	// +kubebuilder:validation:Optional
	ShouldKeep *bool `json:"shouldKeep,omitempty" tf:"should_keep,omitempty"`

	// The rank assigned to the set of answers that match the expression in answerCondition. Answers with the lowest values move to the beginning of the list without changing the relative order of those with the same value. Answers can be given a value between 0 and 255.
	// +kubebuilder:validation:Optional
	Value *float64 `json:"value,omitempty" tf:"value,omitempty"`
}

type SteeringPolicyInitParameters struct {

	// The set of all answers that can potentially issue from the steering policy.
	Answers []AnswersInitParameters `json:"answers,omitempty" tf:"answers,omitempty"`

	// (Updatable) The OCID of the compartment containing the steering policy.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name for the steering policy. Does not have to be unique and can be changed. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The OCID of the health check monitor providing health data about the answers of the steering policy. A steering policy answer with rdata matching a monitored endpoint will use the health data of that endpoint. A steering policy answer with rdata not matching any monitored endpoint will be assumed healthy.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/healthchecks/v1alpha1.HTTPMonitor
	HealthCheckMonitorID *string `json:"healthCheckMonitorId,omitempty" tf:"health_check_monitor_id,omitempty"`

	// Reference to a HTTPMonitor in healthchecks to populate healthCheckMonitorId.
	// +kubebuilder:validation:Optional
	HealthCheckMonitorIDRef *v1.Reference `json:"healthCheckMonitorIdRef,omitempty" tf:"-"`

	// Selector for a HTTPMonitor in healthchecks to populate healthCheckMonitorId.
	// +kubebuilder:validation:Optional
	HealthCheckMonitorIDSelector *v1.Selector `json:"healthCheckMonitorIdSelector,omitempty" tf:"-"`

	// The series of rules that will be processed in sequence to reduce the pool of answers to a response for any given request.
	Rules []SteeringPolicyRulesInitParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// (Updatable) The Time To Live (TTL) for responses from the steering policy, in seconds. If not specified during creation, a value of 30 seconds will be used.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// (Updatable) A set of predefined rules based on the desired purpose of the steering policy. Each template utilizes Traffic Management's rules in a different order to produce the desired results when answering DNS queries.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type SteeringPolicyObservation struct {

	// The set of all answers that can potentially issue from the steering policy.
	Answers []AnswersObservation `json:"answers,omitempty" tf:"answers,omitempty"`

	// (Updatable) The OCID of the compartment containing the steering policy.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name for the steering policy. Does not have to be unique and can be changed. Avoid entering confidential information.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The OCID of the health check monitor providing health data about the answers of the steering policy. A steering policy answer with rdata matching a monitored endpoint will use the health data of that endpoint. A steering policy answer with rdata not matching any monitored endpoint will be assumed healthy.
	HealthCheckMonitorID *string `json:"healthCheckMonitorId,omitempty" tf:"health_check_monitor_id,omitempty"`

	// The OCID of the resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The series of rules that will be processed in sequence to reduce the pool of answers to a response for any given request.
	Rules []SteeringPolicyRulesObservation `json:"rules,omitempty" tf:"rules,omitempty"`

	// The canonical absolute URL of the resource.
	Self *string `json:"self,omitempty" tf:"self,omitempty"`

	// The current state of the resource.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// (Updatable) The Time To Live (TTL) for responses from the steering policy, in seconds. If not specified during creation, a value of 30 seconds will be used.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// (Updatable) A set of predefined rules based on the desired purpose of the steering policy. Each template utilizes Traffic Management's rules in a different order to produce the desired results when answering DNS queries.
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// The date and time the resource was created, expressed in RFC 3339 timestamp format.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type SteeringPolicyParameters struct {

	// The set of all answers that can potentially issue from the steering policy.
	// +kubebuilder:validation:Optional
	Answers []AnswersParameters `json:"answers,omitempty" tf:"answers,omitempty"`

	// (Updatable) The OCID of the compartment containing the steering policy.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name for the steering policy. Does not have to be unique and can be changed. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) The OCID of the health check monitor providing health data about the answers of the steering policy. A steering policy answer with rdata matching a monitored endpoint will use the health data of that endpoint. A steering policy answer with rdata not matching any monitored endpoint will be assumed healthy.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/healthchecks/v1alpha1.HTTPMonitor
	// +kubebuilder:validation:Optional
	HealthCheckMonitorID *string `json:"healthCheckMonitorId,omitempty" tf:"health_check_monitor_id,omitempty"`

	// Reference to a HTTPMonitor in healthchecks to populate healthCheckMonitorId.
	// +kubebuilder:validation:Optional
	HealthCheckMonitorIDRef *v1.Reference `json:"healthCheckMonitorIdRef,omitempty" tf:"-"`

	// Selector for a HTTPMonitor in healthchecks to populate healthCheckMonitorId.
	// +kubebuilder:validation:Optional
	HealthCheckMonitorIDSelector *v1.Selector `json:"healthCheckMonitorIdSelector,omitempty" tf:"-"`

	// The series of rules that will be processed in sequence to reduce the pool of answers to a response for any given request.
	// +kubebuilder:validation:Optional
	Rules []SteeringPolicyRulesParameters `json:"rules,omitempty" tf:"rules,omitempty"`

	// (Updatable) The Time To Live (TTL) for responses from the steering policy, in seconds. If not specified during creation, a value of 30 seconds will be used.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// (Updatable) A set of predefined rules based on the desired purpose of the steering policy. Each template utilizes Traffic Management's rules in a different order to produce the desired results when answering DNS queries.
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`
}

type SteeringPolicyRulesInitParameters struct {

	// An array of caseConditions. A rule may optionally include a sequence of cases defining alternate configurations for how it should behave during processing for any given DNS query. When a rule has no sequence of cases, it is always evaluated with the same configuration during processing. When a rule has an empty sequence of cases, it is always ignored during processing. When a rule has a non-empty sequence of cases, its behavior during processing is configured by the first matching case in the sequence. When a rule has no matching cases the rule is ignored. A rule case with no caseCondition always matches. A rule case with a caseCondition matches only when that expression evaluates to true for the given query.
	Cases []CasesInitParameters `json:"cases,omitempty" tf:"cases,omitempty"`

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) Defines a default set of answer conditions and values that are applied to an answer when cases is not defined for the rule, or a matching case does not have any matching answerConditions in its answerData. defaultAnswerData is not applied if cases is defined and there are no matching cases. In this scenario, the next rule will be processed.
	DefaultAnswerData []DefaultAnswerDataInitParameters `json:"defaultAnswerData,omitempty" tf:"default_answer_data,omitempty"`

	// (Applicable when rule_type=LIMIT) Defines a default count if cases is not defined for the rule or a matching case does not define count. defaultCount is not applied if cases is defined and there are no matching cases. In this scenario, the next rule will be processed. If no rules remain to be processed, the answer will be chosen from the remaining list of answers.
	DefaultCount *float64 `json:"defaultCount,omitempty" tf:"default_count,omitempty"`

	// A user-defined description of the rule's purpose or behavior.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The type of a rule determines its sorting/filtering behavior.
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`
}

type SteeringPolicyRulesObservation struct {

	// An array of caseConditions. A rule may optionally include a sequence of cases defining alternate configurations for how it should behave during processing for any given DNS query. When a rule has no sequence of cases, it is always evaluated with the same configuration during processing. When a rule has an empty sequence of cases, it is always ignored during processing. When a rule has a non-empty sequence of cases, its behavior during processing is configured by the first matching case in the sequence. When a rule has no matching cases the rule is ignored. A rule case with no caseCondition always matches. A rule case with a caseCondition matches only when that expression evaluates to true for the given query.
	Cases []CasesObservation `json:"cases,omitempty" tf:"cases,omitempty"`

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) Defines a default set of answer conditions and values that are applied to an answer when cases is not defined for the rule, or a matching case does not have any matching answerConditions in its answerData. defaultAnswerData is not applied if cases is defined and there are no matching cases. In this scenario, the next rule will be processed.
	DefaultAnswerData []DefaultAnswerDataObservation `json:"defaultAnswerData,omitempty" tf:"default_answer_data,omitempty"`

	// (Applicable when rule_type=LIMIT) Defines a default count if cases is not defined for the rule or a matching case does not define count. defaultCount is not applied if cases is defined and there are no matching cases. In this scenario, the next rule will be processed. If no rules remain to be processed, the answer will be chosen from the remaining list of answers.
	DefaultCount *float64 `json:"defaultCount,omitempty" tf:"default_count,omitempty"`

	// A user-defined description of the rule's purpose or behavior.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The type of a rule determines its sorting/filtering behavior.
	RuleType *string `json:"ruleType,omitempty" tf:"rule_type,omitempty"`
}

type SteeringPolicyRulesParameters struct {

	// An array of caseConditions. A rule may optionally include a sequence of cases defining alternate configurations for how it should behave during processing for any given DNS query. When a rule has no sequence of cases, it is always evaluated with the same configuration during processing. When a rule has an empty sequence of cases, it is always ignored during processing. When a rule has a non-empty sequence of cases, its behavior during processing is configured by the first matching case in the sequence. When a rule has no matching cases the rule is ignored. A rule case with no caseCondition always matches. A rule case with a caseCondition matches only when that expression evaluates to true for the given query.
	// +kubebuilder:validation:Optional
	Cases []CasesParameters `json:"cases,omitempty" tf:"cases,omitempty"`

	// (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) Defines a default set of answer conditions and values that are applied to an answer when cases is not defined for the rule, or a matching case does not have any matching answerConditions in its answerData. defaultAnswerData is not applied if cases is defined and there are no matching cases. In this scenario, the next rule will be processed.
	// +kubebuilder:validation:Optional
	DefaultAnswerData []DefaultAnswerDataParameters `json:"defaultAnswerData,omitempty" tf:"default_answer_data,omitempty"`

	// (Applicable when rule_type=LIMIT) Defines a default count if cases is not defined for the rule or a matching case does not define count. defaultCount is not applied if cases is defined and there are no matching cases. In this scenario, the next rule will be processed. If no rules remain to be processed, the answer will be chosen from the remaining list of answers.
	// +kubebuilder:validation:Optional
	DefaultCount *float64 `json:"defaultCount,omitempty" tf:"default_count,omitempty"`

	// A user-defined description of the rule's purpose or behavior.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The type of a rule determines its sorting/filtering behavior.
	// +kubebuilder:validation:Optional
	RuleType *string `json:"ruleType" tf:"rule_type,omitempty"`
}

// SteeringPolicySpec defines the desired state of SteeringPolicy
type SteeringPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SteeringPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SteeringPolicyInitParameters `json:"initProvider,omitempty"`
}

// SteeringPolicyStatus defines the observed state of SteeringPolicy.
type SteeringPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SteeringPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SteeringPolicy is the Schema for the SteeringPolicys API. Provides the Steering Policy resource in Oracle Cloud Infrastructure DNS service
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type SteeringPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.template) || (has(self.initProvider) && has(self.initProvider.template))",message="spec.forProvider.template is a required parameter"
	Spec   SteeringPolicySpec   `json:"spec"`
	Status SteeringPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SteeringPolicyList contains a list of SteeringPolicys
type SteeringPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SteeringPolicy `json:"items"`
}

// Repository type metadata.
var (
	SteeringPolicy_Kind             = "SteeringPolicy"
	SteeringPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SteeringPolicy_Kind}.String()
	SteeringPolicy_KindAPIVersion   = SteeringPolicy_Kind + "." + CRDGroupVersion.String()
	SteeringPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SteeringPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SteeringPolicy{}, &SteeringPolicyList{})
}
