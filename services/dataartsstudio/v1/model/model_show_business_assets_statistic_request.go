package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBusinessAssetsStatisticRequest Request Object
type ShowBusinessAssetsStatisticRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 分页参数，查询偏移量，默认查询所有
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数，每页数量，默认查询所有
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowBusinessAssetsStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBusinessAssetsStatisticRequest struct{}"
	}

	return strings.Join([]string{"ShowBusinessAssetsStatisticRequest", string(data)}, " ")
}
