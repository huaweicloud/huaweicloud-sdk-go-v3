package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReferenceInfo 脚本版本或文件包版本引用的作业详情。
type ReferenceInfo struct {

	// 作业ID。
	JobId string `json:"job_id"`

	// 作业名称。
	JobName string `json:"job_name"`
}

func (o ReferenceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReferenceInfo struct{}"
	}

	return strings.Join([]string{"ReferenceInfo", string(data)}, " ")
}
