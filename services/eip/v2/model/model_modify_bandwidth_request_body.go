package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyBandwidthRequestBody 批量更新带宽对象的请求体
type ModifyBandwidthRequestBody struct {

	// 更新带宽列表
	Bandwidths []ModifyBandwidthOption `json:"bandwidths"`
}

func (o ModifyBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyBandwidthRequestBody", string(data)}, " ")
}
