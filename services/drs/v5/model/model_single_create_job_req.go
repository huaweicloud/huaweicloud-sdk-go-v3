package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SingleCreateJobReq 创建单个任务请求体。
type SingleCreateJobReq struct {
	Job *CreateJobReq `json:"job"`
}

func (o SingleCreateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleCreateJobReq struct{}"
	}

	return strings.Join([]string{"SingleCreateJobReq", string(data)}, " ")
}
