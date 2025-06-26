package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiagnosisSummaryItem struct {

	// 被诊断实例的ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 被诊断实例的名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 被诊断实例的诊断进度
	Progress *int32 `json:"progress,omitempty"`

	// 诊断任务状态，waiting,executing,failed,finish,cancel
	Status *string `json:"status,omitempty"`

	// 实例的正常诊断项数量
	NormalItemNum *int32 `json:"normal_item_num,omitempty"`

	// 实例的异常诊断项数量
	AbnormalItemNum *int32 `json:"abnormal_item_num,omitempty"`
}

func (o DiagnosisSummaryItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiagnosisSummaryItem struct{}"
	}

	return strings.Join([]string{"DiagnosisSummaryItem", string(data)}, " ")
}
