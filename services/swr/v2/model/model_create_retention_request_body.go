package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRetentionRequestBody struct {

	// 回收规则匹配策略，固定为\"or\"
	Algorithm string `json:"algorithm"`

	// 镜像老化规则
	Rules []Rule `json:"rules"`
}

func (o CreateRetentionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetentionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRetentionRequestBody", string(data)}, " ")
}
