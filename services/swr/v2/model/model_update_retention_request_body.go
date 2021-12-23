package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRetentionRequestBody struct {
	// 算法

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
