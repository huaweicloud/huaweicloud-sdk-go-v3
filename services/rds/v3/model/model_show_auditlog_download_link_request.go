package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowAuditlogDownloadLinkRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *GenerateAuditlogDownloadLinkRequest `json:"body,omitempty"`
}

func (o ShowAuditlogDownloadLinkRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAuditlogDownloadLinkRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditlogDownloadLinkRequest", string(data)}, " ")
}
