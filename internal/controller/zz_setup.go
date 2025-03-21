// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	account "github.com/sibtaina/provider-cloudflare-test/internal/controller/account/account"
	member "github.com/sibtaina/provider-cloudflare-test/internal/controller/account/member"
	providerconfig "github.com/sibtaina/provider-cloudflare-test/internal/controller/providerconfig"
	zone "github.com/sibtaina/provider-cloudflare-test/internal/controller/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		account.Setup,
		member.Setup,
		providerconfig.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
