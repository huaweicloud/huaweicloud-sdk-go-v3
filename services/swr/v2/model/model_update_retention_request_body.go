package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRetentionRequestBody struct {

	// 老化规则匹配策略，固定为\"or\"
	Algorithm string `json:"algorithm"`

	// 镜像老化规则
	Rules []Rule `json:"rules"`
}

func (o UpdateRetentionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRetentionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRetentionRequestBody", string(data)}, " ")
}
