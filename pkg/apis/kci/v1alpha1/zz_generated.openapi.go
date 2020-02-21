// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.Database":         schema_pkg_apis_kci_v1alpha1_Database(ref),
		"github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DatabaseSpec":     schema_pkg_apis_kci_v1alpha1_DatabaseSpec(ref),
		"github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DatabaseStatus":   schema_pkg_apis_kci_v1alpha1_DatabaseStatus(ref),
		"github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DbInstance":       schema_pkg_apis_kci_v1alpha1_DbInstance(ref),
		"github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DbInstanceSpec":   schema_pkg_apis_kci_v1alpha1_DbInstanceSpec(ref),
		"github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DbInstanceStatus": schema_pkg_apis_kci_v1alpha1_DbInstanceStatus(ref),
	}
}

func schema_pkg_apis_kci_v1alpha1_Database(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Database is the Schema for the databases API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DatabaseSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DatabaseStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DatabaseSpec", "github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DatabaseStatus"},
	}
}

func schema_pkg_apis_kci_v1alpha1_DatabaseSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatabaseSpec defines the desired state of Database",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_kci_v1alpha1_DatabaseStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatabaseStatus defines the observed state of Database",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_kci_v1alpha1_DbInstance(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DbInstance is the Schema for the dbinstances API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DbInstanceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DbInstanceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DbInstanceSpec", "github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DbInstanceStatus"},
	}
}

func schema_pkg_apis_kci_v1alpha1_DbInstanceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DbInstanceSpec defines the desired state of DbInstance",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"engine": {
						SchemaProps: spec.SchemaProps{
							Description: "Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"adminUserSecret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"backup": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DbInstanceBackup"),
						},
					},
					"configMap": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"google": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.GoogleInstance"),
						},
					},
					"generic": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.GenericInstance"),
						},
					},
				},
				Required: []string{"engine", "adminUserSecret", "backup"},
			},
		},
		Dependencies: []string{
			"github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.DbInstanceBackup", "github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.GenericInstance", "github.com/kloeckner-i/db-operator/pkg/apis/kci/v1alpha1.GoogleInstance"},
	}
}

func schema_pkg_apis_kci_v1alpha1_DbInstanceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DbInstanceStatus defines the observed state of DbInstance",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"info": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"checksums": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"phase", "status"},
			},
		},
	}
}
