package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeResult 数据库补丁升级结果
type UpgradeResult struct {

	// 任务ID。仅当补丁版本升级任务提交成功时返回该字段。
	JobId *string `json:"job_id,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 错误码。仅当补丁版本升级任务提交失败时返回该字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 失败原因。仅当补丁版本升级任务提交失败时返回该字段。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o UpgradeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeResult struct{}"
	}

	return strings.Join([]string{"UpgradeResult", string(data)}, " ")
}
