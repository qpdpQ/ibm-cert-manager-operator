/*
Copyright 2022.

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

package resources

import (
	certmanagerv1 "github.com/ibm/ibm-cert-manager-operator/apis/cert-manager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Issuer = certmanagerv1.Issuer{
	TypeMeta: metav1.TypeMeta{
		Kind:       "Issuer",
		APIVersion: "cert-manager.io/v1",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:      "smoke-check-issuer",
		Namespace: DeployNamespace,
	},
	Spec: certmanagerv1.IssuerSpec{
		IssuerConfig: certmanagerv1.IssuerConfig{
			SelfSigned: &certmanagerv1.SelfSignedIssuer{},
		},
	},
}
