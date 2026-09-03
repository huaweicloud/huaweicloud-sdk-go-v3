package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestsuiteInfoUsingRequest Request Object
type ShowTestsuiteInfoUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	// 任务id
	SuiteId string `json:"suite_id"`

	// 测试计划Id
	PlanId *string `json:"planId,omitempty"`
}

func (o ShowTestsuiteInfoUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestsuiteInfoUsingRequest struct{}"
	}

	return strings.Join([]string{"ShowTestsuiteInfoUsingRequest", string(data)}, " ")
}
