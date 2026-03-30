// Package check defines the interface for host check plugins.
package check

import (
	"context"
	"time"
)

// Status represents the result status of a check.
type Status string

const (
	StatusPass    Status = "PASS"
	StatusFail    Status = "FAIL"
	StatusPartial Status = "PARTIAL"
	StatusWarn    Status = "WARN"
	StatusError   Status = "ERROR"
	StatusSkipped Status = "SKIPPED"
)

type PluginType string

const (
	// PluginTypeDomainFilter plugins are used to filter hosts based on domain patterns.
	PluginTypeDomainFilter PluginType = "DOMAIN_FILTER"

	// PluginTypeHostCheck plugins perform checks against hosts and return results.
	PluginTypeHostCheck PluginType = "HOST_CHECK"

	// PluginTypeHostDiscovery plugins discover additional hosts related to the input host.
	PluginTypeHostDiscovery PluginType = "HOST_DISCOVERY"

	// PluginTypeSummary plugins generate summary reports based on the results of other checks.
	PluginTypeSummary PluginType = "SUMMARY"
)

var PluginTypeOrder = []PluginType{
	PluginTypeDomainFilter,
	PluginTypeHostCheck,
	PluginTypeHostDiscovery,
	PluginTypeSummary,
}

const DefaultTimeout = 60 * time.Second

type ResultTask struct {
	CheckName string `json:"check_name"`
	Status    Status `json:"status"`
	Message   string `json:"message"`
}

// Result represents the outcome of a check.
type Result struct {
	Name            string       `json:"name"`
	Status          Status       `json:"status"`
	Message         string       `json:"message"`
	Details         []string     `json:"details,omitempty"`
	Duration        string       `json:"duration"`
	Tasks           []ResultTask `json:"tasks,omitempty"`
	AdditionalHosts []string     `json:"additional_hosts,omitempty"`
	HostList        []string     `json:"host_list,omitempty"`
}

type PluginInfo struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Type        PluginType `json:"type"`
}

// Check is the interface that all plugins must implement.
type Check interface {
	// Name returns the unique name of the check.
	Name() string

	// Info returns the plugin information.
	Info() PluginInfo

	// Description returns a human-readable description.
	Description() string

	// Run executes the check against the given hostname.
	Run(ctx context.Context, hostname string, cfg map[string]any, data []Result) Result

	// Version returns the JSON encoded version information of the check (cliversion format).
	Version() []byte
}
