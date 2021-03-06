/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	APIGroup = "ibmclouddatabasesv5.ibm-cloud.crossplane.io"
	Version  = "v1alpha1"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: APIGroup, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// Ibmclouddatabasesv5 types metadata.
var (
	ScalingGroupKind                 = reflect.TypeOf(ScalingGroup{}).Name()
	ScalingGroupGroupKind            = schema.GroupKind{Group: APIGroup, Kind: ScalingGroupKind}.String()
	ScalingGroupKindAPIVersion       = ScalingGroupKind + "." + SchemeGroupVersion.String()
	ScalingGroupGroupVersionKind     = SchemeGroupVersion.WithKind(ScalingGroupKind)
	WhitelistKind                    = reflect.TypeOf(Whitelist{}).Name()
	WhitelistGroupKind               = schema.GroupKind{Group: APIGroup, Kind: WhitelistKind}.String()
	WhitelistKindAPIVersion          = WhitelistKind + "." + SchemeGroupVersion.String()
	WhitelistGroupVersionKind        = SchemeGroupVersion.WithKind(WhitelistKind)
	AutoscalingGroupKind             = reflect.TypeOf(AutoscalingGroup{}).Name()
	AutoscalingGroupGroupKind        = schema.GroupKind{Group: APIGroup, Kind: AutoscalingGroupKind}.String()
	AutoscalingGroupKindAPIVersion   = AutoscalingGroupKind + "." + SchemeGroupVersion.String()
	AutoscalingGroupGroupVersionKind = SchemeGroupVersion.WithKind(AutoscalingGroupKind)
)

func init() {
	SchemeBuilder.Register(
		&ScalingGroup{},
		&ScalingGroupList{},
		&Whitelist{},
		&WhitelistList{},
		&AutoscalingGroup{},
		&AutoscalingGroupList{},
	)
}
