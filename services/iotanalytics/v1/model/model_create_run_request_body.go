package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRunRequestBody struct {

	// 计算资源ID。
	ComputingResourceId string `json:"computing_resource_id"`

	// 作业配置项。
	Conf *[]string `json:"conf,omitempty"`
}

func (o CreateRunRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRunRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRunRequestBody", string(data)}, " ")
}
