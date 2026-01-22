package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchDeleteInstanceReq struct {

	// **参数解释**： 实例的ID列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Instances *[]string `json:"instances,omitempty"`

	// **参数解释**： 对实例的操作：delete。 **约束限制**： 不涉及。 **取值范围**：  - delete **默认取值**： 不涉及
	Action BatchDeleteInstanceReqAction `json:"action"`

	// **参数解释**： 参数值为reliability，表示删除租户所有创建失败的RocketMQ实例。 **约束限制**： 不涉及。 **取值范围**： - reliability **默认取值**： 不涉及。
	AllFailure *BatchDeleteInstanceReqAllFailure `json:"all_failure,omitempty"`

	// **参数解释**： 是否强删除。 **约束限制**： 不涉及。 **取值范围**： - true：强删除，强删除实例不进入回收站。 - false：弱删除，开启回收站功能后，实例进入回收站。 **默认取值**： 不涉及。
	ForceDelete *bool `json:"force_delete,omitempty"`
}

func (o BatchDeleteInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceReq", string(data)}, " ")
}

type BatchDeleteInstanceReqAction struct {
	value string
}

type BatchDeleteInstanceReqActionEnum struct {
	DELETE BatchDeleteInstanceReqAction
}

func GetBatchDeleteInstanceReqActionEnum() BatchDeleteInstanceReqActionEnum {
	return BatchDeleteInstanceReqActionEnum{
		DELETE: BatchDeleteInstanceReqAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteInstanceReqAction) Value() string {
	return c.value
}

func (c BatchDeleteInstanceReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteInstanceReqAction) UnmarshalJSON(b []byte) error {
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

type BatchDeleteInstanceReqAllFailure struct {
	value string
}

type BatchDeleteInstanceReqAllFailureEnum struct {
	RELIABILITY BatchDeleteInstanceReqAllFailure
}

func GetBatchDeleteInstanceReqAllFailureEnum() BatchDeleteInstanceReqAllFailureEnum {
	return BatchDeleteInstanceReqAllFailureEnum{
		RELIABILITY: BatchDeleteInstanceReqAllFailure{
			value: "reliability",
		},
	}
}

func (c BatchDeleteInstanceReqAllFailure) Value() string {
	return c.value
}

func (c BatchDeleteInstanceReqAllFailure) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteInstanceReqAllFailure) UnmarshalJSON(b []byte) error {
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
