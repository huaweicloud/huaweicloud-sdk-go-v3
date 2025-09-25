package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFeatureRequest Request Object
type ShowFeatureRequest struct {

	// 特性key, 当前支持： - replication.enable - replication.source_region - replication.destination_regions - replication.destination_dgw_regions - features.backup_double_az
	FeatureKey string `json:"feature_key"`

	// 每页显示条目数，正整数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移值,正整数
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowFeatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFeatureRequest struct{}"
	}

	return strings.Join([]string{"ShowFeatureRequest", string(data)}, " ")
}
