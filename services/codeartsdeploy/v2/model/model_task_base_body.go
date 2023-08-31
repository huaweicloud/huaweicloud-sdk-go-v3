package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskBaseBody struct {

	// 创建的部署任务id
	Id *string `json:"id,omitempty"`
}

func (o TaskBaseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskBaseBody struct{}"
	}

	return strings.Join([]string{"TaskBaseBody", string(data)}, " ")
}
