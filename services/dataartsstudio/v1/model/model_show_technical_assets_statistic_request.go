package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTechnicalAssetsStatisticRequest Request Object
type ShowTechnicalAssetsStatisticRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 标签名，指定标签名称可以获取此标签对应技术资产的统计信息。
	Tag *string `json:"tag,omitempty"`

	// 分页参数，查询偏移量，默认查询所有
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数，每页数量，默认查询所有
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowTechnicalAssetsStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTechnicalAssetsStatisticRequest struct{}"
	}

	return strings.Join([]string{"ShowTechnicalAssetsStatisticRequest", string(data)}, " ")
}
