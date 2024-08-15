package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstancesVersionResponse Response Object
type UpgradeInstancesVersionResponse struct {

	// 任务id。
	JobIds *[]string `json:"job_ids,omitempty"`

	// 下发成功的实例数量。
	SucceededNum *int32 `json:"succeeded_num,omitempty"`

	// 下发失败的实例数量。
	FailedNum *int32 `json:"failed_num,omitempty"`

	// 下发失败的实例ID列表。
	FailedInstanceIds *[]string `json:"failed_instance_ids,omitempty"`

	// 下发失败错误信息列表。
	ErrorMessages  *[]string `json:"error_messages,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpgradeInstancesVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstancesVersionResponse struct{}"
	}

	return strings.Join([]string{"UpgradeInstancesVersionResponse", string(data)}, " ")
}
