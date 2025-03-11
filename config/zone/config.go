package zone

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_zone", func(r *config.Resource) {
		r.ShortGroup = "zone"
		// r.References["account_id"] = config.Reference{
		// 	Type: "github.com/cdloh/provider-cloudflare/apis/account/v1alpha1.Account",
		// }
	})
}
