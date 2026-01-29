package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectConfigRequest Request Object
type CreateCollectConfigRequest struct {

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	Body *CreateCollectConfigV2RequestBody `json:"body,omitempty"`
}

func (o CreateCollectConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateCollectConfigRequest", string(data)}, " ")
}
