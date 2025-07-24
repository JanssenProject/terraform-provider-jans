
package jans

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// LogPagedResult represents the response structure for audit logs
type LogPagedResult struct {
	Start             int      `json:"start,omitempty"`
	TotalEntriesCount int      `json:"totalEntriesCount,omitempty"`
	EntriesCount      int      `json:"entriesCount,omitempty"`
	Entries           []string `json:"entries,omitempty"`
}

// GetAuditLogs retrieves audit log entries
func (c *Client) GetAuditLogs(ctx context.Context, pattern string, startIndex, limit int, startDate, endDate string) (*LogPagedResult, error) {
	token, err := c.getToken(ctx, "https://jans.io/oauth/config/audit-read")
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	values := url.Values{}
	if pattern != "" {
		values.Set("pattern", pattern)
	}
	if startIndex > 0 {
		values.Set("startIndex", strconv.Itoa(startIndex))
	}
	if limit > 0 {
		values.Set("limit", strconv.Itoa(limit))
	}
	if startDate != "" {
		values.Set("start_date", startDate)
	}
	if endDate != "" {
		values.Set("end_date", endDate)
	}

	endpoint := "/jans-config-api/jans-config-api/api/v1/audit"
	if len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	ret := &LogPagedResult{}
	err = c.get(ctx, endpoint, token, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
