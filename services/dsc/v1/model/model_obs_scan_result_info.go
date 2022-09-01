package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObsScanResultInfo struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// OBS桶ID
	BucketId *string `json:"bucket_id,omitempty" xml:"bucket_id"`

	// OBS桶名称
	BucketName *string `json:"bucket_name,omitempty" xml:"bucket_name"`

	// 文件路径
	FilePath *string `json:"file_path,omitempty" xml:"file_path"`

	// 文件名
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	// 文件md5值
	Md5 *string `json:"md5,omitempty" xml:"md5"`

	// 风险等级
	RiskLevel *int32 `json:"risk_level,omitempty" xml:"risk_level"`

	// 风险数据类型
	SensitiveDataType *[]string `json:"sensitive_data_type,omitempty" xml:"sensitive_data_type"`
}

func (o ObsScanResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsScanResultInfo struct{}"
	}

	return strings.Join([]string{"ObsScanResultInfo", string(data)}, " ")
}
