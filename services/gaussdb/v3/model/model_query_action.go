package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QueryAction struct {

	// 实例或节点动作ID。
	Id *string `json:"id,omitempty"`

	// 实例或节点动作名称。
	Action *string `json:"action,omitempty"`

	// 实例或节点动作对象ID。
	ObjectId *string `json:"object_id,omitempty"`

	// 实例或节点动作类型。
	Type *string `json:"type,omitempty"`

	// 实例或节点动作任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 实例或节点动作状态。
	Status *QueryActionStatus `json:"status,omitempty"`

	// 实例或节点动作创建时间。
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 实例或节点动作更新时间。
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

func (o QueryAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAction struct{}"
	}

	return strings.Join([]string{"QueryAction", string(data)}, " ")
}

type QueryActionStatus struct {
	value string
}

type QueryActionStatusEnum struct {
	OK_TO_RUN QueryActionStatus
	DELETED   QueryActionStatus
}

func GetQueryActionStatusEnum() QueryActionStatusEnum {
	return QueryActionStatusEnum{
		OK_TO_RUN: QueryActionStatus{
			value: "OK_TO_RUN",
		},
		DELETED: QueryActionStatus{
			value: "DELETED",
		},
	}
}

func (c QueryActionStatus) Value() string {
	return c.value
}

func (c QueryActionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryActionStatus) UnmarshalJSON(b []byte) error {
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
