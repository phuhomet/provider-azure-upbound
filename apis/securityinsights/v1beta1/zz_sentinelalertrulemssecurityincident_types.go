/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SentinelAlertRuleMSSecurityIncidentObservation struct {

	// The GUID of the alert rule template which is used to create this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// The description of this Sentinel MS Security Incident Alert Rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Only create incidents when the alert display name doesn't contain text from this list.
	DisplayNameExcludeFilter []*string `json:"displayNameExcludeFilter,omitempty" tf:"display_name_exclude_filter,omitempty"`

	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	DisplayNameFilter []*string `json:"displayNameFilter,omitempty" tf:"display_name_filter,omitempty"`

	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Sentinel MS Security Incident Alert Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// The Microsoft Security Service from where the alert will be generated. Possible values are Azure Active Directory Identity Protection, Azure Advanced Threat Protection, Azure Security Center, Azure Security Center for IoT, Microsoft Cloud App Security, Microsoft Defender Advanced Threat Protection and Office 365 Advanced Threat Protection.
	ProductFilter *string `json:"productFilter,omitempty" tf:"product_filter,omitempty"`

	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are High, Medium, Low and Informational.
	SeverityFilter []*string `json:"severityFilter,omitempty" tf:"severity_filter,omitempty"`
}

type SentinelAlertRuleMSSecurityIncidentParameters struct {

	// The GUID of the alert rule template which is used to create this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	// +kubebuilder:validation:Optional
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// The description of this Sentinel MS Security Incident Alert Rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Only create incidents when the alert display name doesn't contain text from this list.
	// +kubebuilder:validation:Optional
	DisplayNameExcludeFilter []*string `json:"displayNameExcludeFilter,omitempty" tf:"display_name_exclude_filter,omitempty"`

	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	// +kubebuilder:validation:Optional
	DisplayNameFilter []*string `json:"displayNameFilter,omitempty" tf:"display_name_filter,omitempty"`

	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationsmanagement/v1beta1.LogAnalyticsSolution
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("workspace_resource_id",false)
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Reference to a LogAnalyticsSolution in operationsmanagement to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDRef *v1.Reference `json:"logAnalyticsWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a LogAnalyticsSolution in operationsmanagement to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDSelector *v1.Selector `json:"logAnalyticsWorkspaceIdSelector,omitempty" tf:"-"`

	// The Microsoft Security Service from where the alert will be generated. Possible values are Azure Active Directory Identity Protection, Azure Advanced Threat Protection, Azure Security Center, Azure Security Center for IoT, Microsoft Cloud App Security, Microsoft Defender Advanced Threat Protection and Office 365 Advanced Threat Protection.
	// +kubebuilder:validation:Optional
	ProductFilter *string `json:"productFilter,omitempty" tf:"product_filter,omitempty"`

	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are High, Medium, Low and Informational.
	// +kubebuilder:validation:Optional
	SeverityFilter []*string `json:"severityFilter,omitempty" tf:"severity_filter,omitempty"`
}

// SentinelAlertRuleMSSecurityIncidentSpec defines the desired state of SentinelAlertRuleMSSecurityIncident
type SentinelAlertRuleMSSecurityIncidentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelAlertRuleMSSecurityIncidentParameters `json:"forProvider"`
}

// SentinelAlertRuleMSSecurityIncidentStatus defines the observed state of SentinelAlertRuleMSSecurityIncident.
type SentinelAlertRuleMSSecurityIncidentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelAlertRuleMSSecurityIncidentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleMSSecurityIncident is the Schema for the SentinelAlertRuleMSSecurityIncidents API. Manages a Sentinel MS Security Incident Alert Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelAlertRuleMSSecurityIncident struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.productFilter)",message="productFilter is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.severityFilter)",message="severityFilter is a required parameter"
	Spec   SentinelAlertRuleMSSecurityIncidentSpec   `json:"spec"`
	Status SentinelAlertRuleMSSecurityIncidentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleMSSecurityIncidentList contains a list of SentinelAlertRuleMSSecurityIncidents
type SentinelAlertRuleMSSecurityIncidentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelAlertRuleMSSecurityIncident `json:"items"`
}

// Repository type metadata.
var (
	SentinelAlertRuleMSSecurityIncident_Kind             = "SentinelAlertRuleMSSecurityIncident"
	SentinelAlertRuleMSSecurityIncident_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SentinelAlertRuleMSSecurityIncident_Kind}.String()
	SentinelAlertRuleMSSecurityIncident_KindAPIVersion   = SentinelAlertRuleMSSecurityIncident_Kind + "." + CRDGroupVersion.String()
	SentinelAlertRuleMSSecurityIncident_GroupVersionKind = CRDGroupVersion.WithKind(SentinelAlertRuleMSSecurityIncident_Kind)
)

func init() {
	SchemeBuilder.Register(&SentinelAlertRuleMSSecurityIncident{}, &SentinelAlertRuleMSSecurityIncidentList{})
}
