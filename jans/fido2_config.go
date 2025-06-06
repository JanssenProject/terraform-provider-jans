package jans

import (
	"context"
	"fmt"
)

type RequestedParties struct {
	Id      string   `schema:"name" json:"name"`
	Origins []string `schema:"domains" json:"domains"`
}

// Fido2Configuration represents the Fido2 configuration properties
type Fido2Configuration struct {
	AuthenticatorCertsFolder        string             `schema:"authenticator_certs_folder" json:"authenticatorCertsFolder"`
	MdsCertsFolder                  string             `schema:"mds_certs_folder" json:"mdsCertsFolder"`
	MdsTocsFolder                   string             `schema:"mds_tocs_folder" json:"mdsTocsFolder"`
	ServerMetadataFolder            string             `schema:"server_metadata_folder" json:"serverMetadataFolder"`
	RequestedParties                []RequestedParties `schema:"requested_parties" json:"requestedParties"`
	UnfinishedRequestExpiration     int                `schema:"unfinished_request_expiration" json:"unfinishedRequestExpiration"`
	AuthenticationHistoryExpiration int                `schema:"authentication_history_expiration" json:"authenticationHistoryExpiration"`
	DebugUserAutoEnrollment         bool               `schema:"user_auto_enrollment" json:"userAutoEnrollment"`
	EnabledFidoAlgorithms           []string           `schema:"requested_credential_types" json:"enabledFidoAlgorithms"`
}

// JansFido2DynConfiguration defines the Fido2 dynamic configuration
// of the Janssen server.
type JansFido2DynConfiguration struct {
	Issuer                      string             `schema:"issuer" json:"issuer,omitempty"`
	BaseEndpoint                string             `schema:"base_endpoint" json:"baseEndpoint,omitempty"`
	UserInfoLifetime            int                `schema:"user_info_lifetime" json:"userInfoLifetime,omitempty"`
	UseLocalCache               bool               `schema:"use_local_cache" json:"useLocalCache,omitempty"`
	DisableJdkLogger            bool               `schema:"disable_jdk_logger" json:"disableJdkLogger,omitempty"`
	LoggingLevel                string             `schema:"logging_level" json:"loggingLevel,omitempty"`
	LoggingLayout               string             `schema:"logging_layout" json:"loggingLayout,omitempty"`
	ExternalLoggerConfiguration string             `schema:"external_logger_configuration" json:"externalLoggerConfiguration,omitempty"`
	MetricReporterEnabled       bool               `schema:"metric_reporter_enabled" json:"metricReporterEnabled,omitempty"`
	MetricReporterInterval      int                `schema:"metric_reporter_interval" json:"metricReporterInterval,omitempty"`
	MetricReporterKeepDataDays  int                `schema:"metric_reporter_keep_data_days" json:"metricReporterKeepDataDays,omitempty"`
	PersonCustomObjectClassList []string           `schema:"person_custom_object_class_list" json:"personCustomObjectClassList,omitempty"`
	Fido2Configuration          Fido2Configuration `schema:"fido2_configuration" json:"fido2Configuration,omitempty"`
}

// GetFido2Configuration returns the current Fido2 configuration.
func (c *Client) GetFido2Configuration(ctx context.Context) (*JansFido2DynConfiguration, error) {

	token, err := c.getToken(ctx, "https://jans.io/oauth/config/fido2.readonly")
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	ret := &JansFido2DynConfiguration{}

	if err := c.get(ctx, "/jans-config-api/fido2/fido2-config", token, ret); err != nil {
		return nil, fmt.Errorf("get request failed: %w", err)
	}

	return ret, nil
}

// UpdateFido2Configuration updates Fido2 configuration for the Janssen server.
func (c *Client) UpdateFido2Configuration(ctx context.Context, fido2Config *JansFido2DynConfiguration) (*JansFido2DynConfiguration, error) {

	token, err := c.getToken(ctx, "https://jans.io/oauth/config/fido2.write")
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	ret := &JansFido2DynConfiguration{}

	if err := c.put(ctx, "/jans-config-api/fido2/fido2-config", token, fido2Config, ret); err != nil {
		return nil, fmt.Errorf("put request failed: %w", err)
	}

	return ret, nil
}
