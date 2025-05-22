package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListOfflineKeyAnalysisTaskRequest Request Object
type ListOfflineKeyAnalysisTaskRequest struct {

	// **参数解释**： 实例ID。可通过DCS控制台进入实例详情界面查看。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 偏移量，表示从此偏移量开始查询。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 每页显示的条目数量。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 离线分析的任务状态。 **约束限制**： 不涉及。 **取值范围**： waiting：任务等待中。 running：任务进行中。 success：任务执行成功。 failed：任务执行失败。 **默认取值**： 不涉及。
	Status *ListOfflineKeyAnalysisTaskRequestStatus `json:"status,omitempty"`
}

func (o ListOfflineKeyAnalysisTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOfflineKeyAnalysisTaskRequest struct{}"
	}

	return strings.Join([]string{"ListOfflineKeyAnalysisTaskRequest", string(data)}, " ")
}

type ListOfflineKeyAnalysisTaskRequestStatus struct {
	value string
}

type ListOfflineKeyAnalysisTaskRequestStatusEnum struct {
	WAITING ListOfflineKeyAnalysisTaskRequestStatus
	RUNNING ListOfflineKeyAnalysisTaskRequestStatus
	SUCCESS ListOfflineKeyAnalysisTaskRequestStatus
	FAILED  ListOfflineKeyAnalysisTaskRequestStatus
}

func GetListOfflineKeyAnalysisTaskRequestStatusEnum() ListOfflineKeyAnalysisTaskRequestStatusEnum {
	return ListOfflineKeyAnalysisTaskRequestStatusEnum{
		WAITING: ListOfflineKeyAnalysisTaskRequestStatus{
			value: "waiting",
		},
		RUNNING: ListOfflineKeyAnalysisTaskRequestStatus{
			value: "running",
		},
		SUCCESS: ListOfflineKeyAnalysisTaskRequestStatus{
			value: "success",
		},
		FAILED: ListOfflineKeyAnalysisTaskRequestStatus{
			value: "failed",
		},
	}
}

func (c ListOfflineKeyAnalysisTaskRequestStatus) Value() string {
	return c.value
}

func (c ListOfflineKeyAnalysisTaskRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOfflineKeyAnalysisTaskRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
