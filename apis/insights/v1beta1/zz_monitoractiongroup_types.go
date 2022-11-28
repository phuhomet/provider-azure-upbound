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

type AADAuthObservation struct {
}

type AADAuthParameters struct {

	// The identifier URI for AAD auth.
	// +kubebuilder:validation:Optional
	IdentifierURI *string `json:"identifierUri,omitempty" tf:"identifier_uri,omitempty"`

	// The webhook application object Id for AAD auth.
	// +kubebuilder:validation:Required
	ObjectID *string `json:"objectId" tf:"object_id,omitempty"`

	// The tenant id for AAD auth.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ArmRoleReceiverObservation struct {
}

type ArmRoleReceiverParameters struct {

	// The name of the ARM role receiver. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The arm role id.
	// +kubebuilder:validation:Required
	RoleID *string `json:"roleId" tf:"role_id,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type AutomationRunBookReceiverObservation struct {
}

type AutomationRunBookReceiverParameters struct {

	// The automation account ID which holds this runbook and authenticates to Azure resources.
	// +kubebuilder:validation:Required
	AutomationAccountID *string `json:"automationAccountId" tf:"automation_account_id,omitempty"`

	// Indicates whether this instance is global runbook.
	// +kubebuilder:validation:Required
	IsGlobalRunBook *bool `json:"isGlobalRunbook" tf:"is_global_runbook,omitempty"`

	// The name of the automation runbook receiver. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The name for this runbook.
	// +kubebuilder:validation:Required
	RunBookName *string `json:"runbookName" tf:"runbook_name,omitempty"`

	// The URI where webhooks should be sent.
	// +kubebuilder:validation:Required
	ServiceURI *string `json:"serviceUri" tf:"service_uri,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`

	// The resource id for webhook linked to this runbook.
	// +kubebuilder:validation:Required
	WebhookResourceID *string `json:"webhookResourceId" tf:"webhook_resource_id,omitempty"`
}

type AzureAppPushReceiverObservation struct {
}

type AzureAppPushReceiverParameters struct {

	// The email address of the user signed into the mobile app who will receive push notifications from this receiver.
	// +kubebuilder:validation:Required
	EmailAddress *string `json:"emailAddress" tf:"email_address,omitempty"`

	// The name of the Azure app push receiver. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type AzureFunctionReceiverObservation struct {
}

type AzureFunctionReceiverParameters struct {

	// The Azure resource ID of the function app.
	// +kubebuilder:validation:Required
	FunctionAppResourceID *string `json:"functionAppResourceId" tf:"function_app_resource_id,omitempty"`

	// The function name in the function app.
	// +kubebuilder:validation:Required
	FunctionName *string `json:"functionName" tf:"function_name,omitempty"`

	// The HTTP trigger url where HTTP request sent to.
	// +kubebuilder:validation:Required
	HTTPTriggerURL *string `json:"httpTriggerUrl" tf:"http_trigger_url,omitempty"`

	// The name of the Azure Function receiver. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type EmailReceiverObservation struct {
}

type EmailReceiverParameters struct {

	// The email address of this receiver.
	// +kubebuilder:validation:Required
	EmailAddress *string `json:"emailAddress" tf:"email_address,omitempty"`

	// The name of the email receiver. Names must be unique (case-insensitive) across all receivers within an action group. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type EventHubReceiverObservation struct {
}

type EventHubReceiverParameters struct {

	// The resource ID of the respective Event Hub.
	// +kubebuilder:validation:Optional
	EventHubID *string `json:"eventHubId,omitempty" tf:"event_hub_id,omitempty"`

	// The name of the specific Event Hub queue.
	// +kubebuilder:validation:Optional
	EventHubName *string `json:"eventHubName,omitempty" tf:"event_hub_name,omitempty"`

	// The namespace name of the Event Hub.
	// +kubebuilder:validation:Optional
	EventHubNamespace *string `json:"eventHubNamespace,omitempty" tf:"event_hub_namespace,omitempty"`

	// The name of the EventHub Receiver, must be unique within action group. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID for the subscription containing this Event Hub. Default to the subscription ID of the Action Group.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The Tenant ID for the subscription containing this Event Hub.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Indicates whether to use common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type ItsmReceiverObservation struct {
}

type ItsmReceiverParameters struct {

	// The unique connection identifier of the ITSM connection.
	// +kubebuilder:validation:Required
	ConnectionID *string `json:"connectionId" tf:"connection_id,omitempty"`

	// The name of the ITSM receiver. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The region of the workspace.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// A JSON blob for the configurations of the ITSM action. CreateMultipleWorkItems option will be part of this blob as well.
	// +kubebuilder:validation:Required
	TicketConfiguration *string `json:"ticketConfiguration" tf:"ticket_configuration,omitempty"`

	// The Azure Log Analytics workspace ID where this connection is defined. Format is <subscription id>|<workspace id>, for example 00000000-0000-0000-0000-000000000000|00000000-0000-0000-0000-000000000000.
	// +kubebuilder:validation:Required
	WorkspaceID *string `json:"workspaceId" tf:"workspace_id,omitempty"`
}

type LogicAppReceiverObservation struct {
}

type LogicAppReceiverParameters struct {

	// The callback url where HTTP request sent to.
	// +kubebuilder:validation:Required
	CallbackURL *string `json:"callbackUrl" tf:"callback_url,omitempty"`

	// The name of the logic app receiver. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The Azure resource ID of the logic app.
	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type MonitorActionGroupObservation struct {

	// The ID of the Action Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MonitorActionGroupParameters struct {

	// One or more arm_role_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	ArmRoleReceiver []ArmRoleReceiverParameters `json:"armRoleReceiver,omitempty" tf:"arm_role_receiver,omitempty"`

	// One or more automation_runbook_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	AutomationRunBookReceiver []AutomationRunBookReceiverParameters `json:"automationRunbookReceiver,omitempty" tf:"automation_runbook_receiver,omitempty"`

	// One or more azure_app_push_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	AzureAppPushReceiver []AzureAppPushReceiverParameters `json:"azureAppPushReceiver,omitempty" tf:"azure_app_push_receiver,omitempty"`

	// One or more azure_function_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	AzureFunctionReceiver []AzureFunctionReceiverParameters `json:"azureFunctionReceiver,omitempty" tf:"azure_function_receiver,omitempty"`

	// One or more email_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	EmailReceiver []EmailReceiverParameters `json:"emailReceiver,omitempty" tf:"email_receiver,omitempty"`

	// Whether this action group is enabled. If an action group is not enabled, then none of its receivers will receive communications. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more event_hub_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	EventHubReceiver []EventHubReceiverParameters `json:"eventHubReceiver,omitempty" tf:"event_hub_receiver,omitempty"`

	// One or more itsm_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	ItsmReceiver []ItsmReceiverParameters `json:"itsmReceiver,omitempty" tf:"itsm_receiver,omitempty"`

	// One or more logic_app_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	LogicAppReceiver []LogicAppReceiverParameters `json:"logicAppReceiver,omitempty" tf:"logic_app_receiver,omitempty"`

	// The name of the resource group in which to create the Action Group instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// One or more sms_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	SMSReceiver []SMSReceiverParameters `json:"smsReceiver,omitempty" tf:"sms_receiver,omitempty"`

	// The short name of the action group. This will be used in SMS messages.
	// +kubebuilder:validation:Required
	ShortName *string `json:"shortName" tf:"short_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more voice_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	VoiceReceiver []VoiceReceiverParameters `json:"voiceReceiver,omitempty" tf:"voice_receiver,omitempty"`

	// One or more webhook_receiver blocks as defined below.
	// +kubebuilder:validation:Optional
	WebhookReceiver []WebhookReceiverParameters `json:"webhookReceiver,omitempty" tf:"webhook_receiver,omitempty"`
}

type SMSReceiverObservation struct {
}

type SMSReceiverParameters struct {

	// The country code of the SMS receiver.
	// +kubebuilder:validation:Required
	CountryCode *string `json:"countryCode" tf:"country_code,omitempty"`

	// The name of the SMS receiver. Names must be unique (case-insensitive) across all receivers within an action group. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The phone number of the SMS receiver.
	// +kubebuilder:validation:Required
	PhoneNumber *string `json:"phoneNumber" tf:"phone_number,omitempty"`
}

type VoiceReceiverObservation struct {
}

type VoiceReceiverParameters struct {

	// The country code of the voice receiver.
	// +kubebuilder:validation:Required
	CountryCode *string `json:"countryCode" tf:"country_code,omitempty"`

	// The name of the voice receiver. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The phone number of the voice receiver.
	// +kubebuilder:validation:Required
	PhoneNumber *string `json:"phoneNumber" tf:"phone_number,omitempty"`
}

type WebhookReceiverObservation struct {
}

type WebhookReceiverParameters struct {

	// The aad_auth block as defined below
	// +kubebuilder:validation:Optional
	AADAuth []AADAuthParameters `json:"aadAuth,omitempty" tf:"aad_auth,omitempty"`

	// The name of the webhook receiver. Names must be unique (case-insensitive) across all receivers within an action group. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The URI where webhooks should be sent.
	// +kubebuilder:validation:Required
	ServiceURI *string `json:"serviceUri" tf:"service_uri,omitempty"`

	// Enables or disables the common alert schema.
	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

// MonitorActionGroupSpec defines the desired state of MonitorActionGroup
type MonitorActionGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorActionGroupParameters `json:"forProvider"`
}

// MonitorActionGroupStatus defines the observed state of MonitorActionGroup.
type MonitorActionGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorActionGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActionGroup is the Schema for the MonitorActionGroups API. Manages an Action Group within Azure Monitor
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorActionGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorActionGroupSpec   `json:"spec"`
	Status            MonitorActionGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActionGroupList contains a list of MonitorActionGroups
type MonitorActionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorActionGroup `json:"items"`
}

// Repository type metadata.
var (
	MonitorActionGroup_Kind             = "MonitorActionGroup"
	MonitorActionGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorActionGroup_Kind}.String()
	MonitorActionGroup_KindAPIVersion   = MonitorActionGroup_Kind + "." + CRDGroupVersion.String()
	MonitorActionGroup_GroupVersionKind = CRDGroupVersion.WithKind(MonitorActionGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorActionGroup{}, &MonitorActionGroupList{})
}
