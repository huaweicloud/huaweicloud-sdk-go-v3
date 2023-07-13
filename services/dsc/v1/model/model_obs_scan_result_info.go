package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObsScanResultInfo struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// OBS桶ID
	BucketId *string `json:"bucket_id,omitempty"`

	// OBS桶名称
	BucketName *string `json:"bucket_name,omitempty"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty"`

	// 文件名
	FileName *string `json:"file_name,omitempty"`

	// 文件md5值
	Md5 *string `json:"md5,omitempty"`

	// 风险等级
	RiskLevel *int32 `json:"risk_level,omitempty"`

	// 风险数据类型
	SensitiveDataType *[]string `json:"sensitive_data_type,omitempty"`
}

func (o ObsScanResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsScanResultInfo struct{}"
	}

	return strings.Join([]string{"ObsScanResultInfo", string(data)}, " ")
}
