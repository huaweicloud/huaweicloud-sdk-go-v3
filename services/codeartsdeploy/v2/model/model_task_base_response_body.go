package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskBaseResponseBody struct {

	// 部署任务id
	Id *string `json:"id,omitempty"`

	// 部署任务状态
	State *string `json:"state,omitempty"`

	// 部署任务类型
	DeploySystem *string `json:"deploy_system,omitempty"`
}

func (o TaskBaseResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskBaseResponseBody struct{}"
	}

	return strings.Join([]string{"TaskBaseResponseBody", string(data)}, " ")
}
