package jans

import (
	"context"
	"fmt"
)

// ScimAppConfigurations is the definition of the SCIM app configuration.
type ScimAppConfigurations struct {
	BaseDN                      string `schema:"base_dn" json:"baseDN"`
	DisableLoggerTimer          bool   `schema:"disable_logger_timer" json:"disableLoggerTimer"`
	DisableAuditLogger          bool   `schema:"disable_audit_logger" json:"disableAuditLogger"`
	ApplicationUrl              string `schema:"application_url" json:"applicationUrl"`
	BaseEndpoint                string `schema:"base_endpoint" json:"baseEndpoint"`
	PersonCustomObjectClass     string `schema:"person_custom_object_class" json:"personCustomObjectClass"`
	OxAuthIssuer                string `schema:"ox_auth_issuer" json:"oxAuthIssuer"`
	ProtectionMode              string `schema:"protection_mode" json:"protectionMode"`
	MaxCount                    int    `schema:"max_count" json:"maxCount"`
	UserExtensionSchemaURI      string `schema:"user_extension_schema_uri" json:"userExtensionSchemaURI"`
	LoggingLevel                string `schema:"logging_level" json:"loggingLevel"`
	LoggingLayout               string `schema:"logging_layout" json:"loggingLayout"`
	ExternalLoggerConfiguration string `schema:"external_logger_configuration" json:"externalLoggerConfiguration"`
	MetricReporterInterval      int    `schema:"metric_reporter_interval" json:"metricReporterInterval"`
	MetricReporterKeepDataDays  int    `schema:"metric_reporter_keep_data_days" json:"metricReporterKeepDataDays"`
	MetricReporterEnabled       bool   `schema:"metric_reporter_enabled" json:"metricReporterEnabled"`
	DisableJdkLogger            bool   `schema:"disable_jdk_logger" json:"disableJdkLogger"`
	UseLocalCache               bool   `schema:"use_local_cache" json:"useLocalCache"`
	BulkMaxOperations           int    `schema:"bulk_max_operations" json:"bulkMaxOperations"`
	BulkMaxPayloadSize          int    `schema:"bulk_max_payload_size" json:"bulkMaxPayloadSize"`
}

// GetScimAppConfiguration returns the current SCIM App configuration.
func (c *Client) GetScimAppConfiguration(ctx context.Context) (*ScimAppConfigurations, error) {

	token, err := c.getToken(ctx, "https://jans.io/scim/config.readonly")
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	ret := &ScimAppConfigurations{}

	if err := c.get(ctx, "/jans-config-api/scim/scim-config", token, ret); err != nil {
		return nil, fmt.Errorf("get request failed: %w", err)
	}

	return ret, nil
}

// PatchScimAppConfiguration updates the SCIM App configuration based on the
// provided set of patches.
func (c *Client) PatchScimAppConfiguration(ctx context.Context, patches []PatchRequest) (*ScimAppConfigurations, error) {

	if len(patches) == 0 {
		return c.GetScimAppConfiguration(ctx)
	}

	orig, err := c.GetScimAppConfiguration(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get scim configuration: %w", err)
	}

	updates, err := createPatchesDiff(orig, patches)
	if err != nil {
		return nil, fmt.Errorf("failed to create patches: %w", err)
	}

	if len(updates) == 0 {
		return c.GetScimAppConfiguration(ctx)
	}

	token, err := c.getToken(ctx, "https://jans.io/scim/config.write")
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	if err := c.patch(ctx, "/jans-config-api/scim/scim-config", token, updates); err != nil {
		return nil, fmt.Errorf("patch request failed: %w", err)
	}

	return c.GetScimAppConfiguration(ctx)
}
