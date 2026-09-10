package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelServiceTaskReq 调用模型服务创建任务请求体
type CreateModelServiceTaskReq struct {
	Inputs *ModelServiceTaskInputs `json:"inputs"`
}

func (o CreateModelServiceTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelServiceTaskReq struct{}"
	}

	return strings.Join([]string{"CreateModelServiceTaskReq", string(data)}, " ")
}
