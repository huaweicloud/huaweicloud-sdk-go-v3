package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ValidationResult struct {

	// 创建或更新连接校验结果，请参见linkConfig参数说明
	LinkConfig *[]ValidationLinkConfig `json:"linkConfig,omitempty"`
}

func (o ValidationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidationResult struct{}"
	}

	return strings.Join([]string{"ValidationResult", string(data)}, " ")
}
