package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTopologyTreeRequest struct {

	// 区域id。
	RegionId *string `json:"region_id,omitempty"`

	// 应用id。
	BusinessId int64 `json:"business_id"`

	// 环境标签id。
	EnvTagId *int64 `json:"env_tag_id,omitempty"`

	// 环境关键字。
	EnvKeyword *string `json:"env_keyword,omitempty"`

	// 应用id。
	XBusinessId int64 `json:"x-business-id"`
}

func (o ShowTopologyTreeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopologyTreeRequest struct{}"
	}

	return strings.Join([]string{"ShowTopologyTreeRequest", string(data)}, " ")
}
