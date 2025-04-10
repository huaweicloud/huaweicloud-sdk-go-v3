package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetDrConfigRequestBody 重置容灾配置request信息。
type ResetDrConfigRequestBody struct {

	// 对端子网IP网段或者IP，多个以英文逗号分割。
	OppositeDataCidr *string `json:"opposite_data_cidr,omitempty"`
}

func (o ResetDrConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetDrConfigRequestBody struct{}"
	}

	return strings.Join([]string{"ResetDrConfigRequestBody", string(data)}, " ")
}
