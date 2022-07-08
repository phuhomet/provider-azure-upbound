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

type ContainerConnectedRegistryObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ContainerConnectedRegistryParameters struct {

	// +kubebuilder:validation:Optional
	AuditLogEnabled *bool `json:"auditLogEnabled,omitempty" tf:"audit_log_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ClientTokenIds []*string `json:"clientTokenIds,omitempty" tf:"client_token_ids,omitempty"`

	// +crossplane:generate:reference:type=Registry
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ContainerRegistryID *string `json:"containerRegistryId,omitempty" tf:"container_registry_id,omitempty"`

	// +kubebuilder:validation:Optional
	ContainerRegistryIDRef *v1.Reference `json:"containerRegistryIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ContainerRegistryIDSelector *v1.Selector `json:"containerRegistryIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	Notification []NotificationParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// +kubebuilder:validation:Optional
	ParentRegistryID *string `json:"parentRegistryId,omitempty" tf:"parent_registry_id,omitempty"`

	// +kubebuilder:validation:Optional
	SyncMessageTTL *string `json:"syncMessageTtl,omitempty" tf:"sync_message_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	SyncSchedule *string `json:"syncSchedule,omitempty" tf:"sync_schedule,omitempty"`

	// +crossplane:generate:reference:type=Token
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SyncTokenID *string `json:"syncTokenId,omitempty" tf:"sync_token_id,omitempty"`

	// +kubebuilder:validation:Optional
	SyncTokenIDRef *v1.Reference `json:"syncTokenIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SyncTokenIDSelector *v1.Selector `json:"syncTokenIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SyncWindow *string `json:"syncWindow,omitempty" tf:"sync_window,omitempty"`
}

type NotificationObservation struct {
}

type NotificationParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	Digest *string `json:"digest,omitempty" tf:"digest,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

// ContainerConnectedRegistrySpec defines the desired state of ContainerConnectedRegistry
type ContainerConnectedRegistrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerConnectedRegistryParameters `json:"forProvider"`
}

// ContainerConnectedRegistryStatus defines the observed state of ContainerConnectedRegistry.
type ContainerConnectedRegistryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerConnectedRegistryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerConnectedRegistry is the Schema for the ContainerConnectedRegistrys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ContainerConnectedRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerConnectedRegistrySpec   `json:"spec"`
	Status            ContainerConnectedRegistryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerConnectedRegistryList contains a list of ContainerConnectedRegistrys
type ContainerConnectedRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerConnectedRegistry `json:"items"`
}

// Repository type metadata.
var (
	ContainerConnectedRegistry_Kind             = "ContainerConnectedRegistry"
	ContainerConnectedRegistry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContainerConnectedRegistry_Kind}.String()
	ContainerConnectedRegistry_KindAPIVersion   = ContainerConnectedRegistry_Kind + "." + CRDGroupVersion.String()
	ContainerConnectedRegistry_GroupVersionKind = CRDGroupVersion.WithKind(ContainerConnectedRegistry_Kind)
)

func init() {
	SchemeBuilder.Register(&ContainerConnectedRegistry{}, &ContainerConnectedRegistryList{})
}
