package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyEngineAttachmentSummary struct {
	EntityType *EntityType `json:"entity_type"`

	// The unique identifier of the attached entity.
	EntityId string `json:"entity_id"`

	Mode *PolicyEngineMode `json:"mode"`

	// Timestamp in RFC 3339 format (UTC)
	AttachedAt *sdktime.SdkTime `json:"attached_at"`
}

func (o PolicyEngineAttachmentSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyEngineAttachmentSummary struct{}"
	}

	return strings.Join([]string{"PolicyEngineAttachmentSummary", string(data)}, " ")
}
