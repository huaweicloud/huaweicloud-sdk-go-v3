package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobReq 更新作业参数请求体
type UpdateJobReq struct {

	// 步骤的参数修改信息
	Tasks []JobTaskDto `json:"tasks"`
}

func (o UpdateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobReq struct{}"
	}

	return strings.Join([]string{"UpdateJobReq", string(data)}, " ")
}
