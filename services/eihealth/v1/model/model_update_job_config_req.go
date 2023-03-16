package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 作业配置请求体
type UpdateJobConfigReq struct {

	// 作业保存条数
	JobRetainNumber int32 `json:"job_retain_number"`
}

func (o UpdateJobConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobConfigReq struct{}"
	}

	return strings.Join([]string{"UpdateJobConfigReq", string(data)}, " ")
}
