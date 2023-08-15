package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyConfigurationRequestBody struct {

	// 实例ID列表。列表长度限制在10以内。
	InstanceIds []string `json:"instance_ids"`
}

func (o ApplyConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationRequestBody", string(data)}, " ")
}
