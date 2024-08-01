package queue

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("solacebroker_msg_vpn_queue", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource
        r.ShortGroup = "queue"
    })
}
