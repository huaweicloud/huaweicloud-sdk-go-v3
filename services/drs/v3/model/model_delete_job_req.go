package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteJobReq 删除在线迁移任务请求体。
type DeleteJobReq struct {

	// **参数解释**：  批量结束或删除任务的操作类型，指定对任务执行结束还是删除操作。  **约束限制**：  包年/包月任务需要先手动退订，然后delete_type参数设置为delete调用该接口删除任务。  **取值范围**：  - terminate：结束任务 - force_terminate：强制结束任务 - delete：删除任务 - roll_back_upgrade_active_standby：转主备任务回滚  **默认取值**：  不涉及
	DeleteType DeleteJobReqDeleteType `json:"delete_type"`

	// **参数解释**：  任务ID，用于唯一标识需要结束或删除的数据复制任务。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	JobId string `json:"job_id"`

	// **参数解释**：  MySQL为源的数据复制任务，结束任务时是否展示断点信息。  **约束限制**：  仅MySQL为源的实时迁移、实时同步、实时灾备任务适用。  **取值范围**：  - true：展示断点信息 - false：不展示断点信息  **默认取值**：  不涉及
	IsShowBreakpointPosition *bool `json:"is_show_breakpoint_position,omitempty"`
}

func (o DeleteJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobReq struct{}"
	}

	return strings.Join([]string{"DeleteJobReq", string(data)}, " ")
}

type DeleteJobReqDeleteType struct {
	value string
}

type DeleteJobReqDeleteTypeEnum struct {
	TERMINATE                        DeleteJobReqDeleteType
	FORCE_TERMINATE                  DeleteJobReqDeleteType
	DELETE                           DeleteJobReqDeleteType
	ROLL_BACK_UPGRADE_ACTIVE_STANDBY DeleteJobReqDeleteType
}

func GetDeleteJobReqDeleteTypeEnum() DeleteJobReqDeleteTypeEnum {
	return DeleteJobReqDeleteTypeEnum{
		TERMINATE: DeleteJobReqDeleteType{
			value: "terminate",
		},
		FORCE_TERMINATE: DeleteJobReqDeleteType{
			value: "force_terminate",
		},
		DELETE: DeleteJobReqDeleteType{
			value: "delete",
		},
		ROLL_BACK_UPGRADE_ACTIVE_STANDBY: DeleteJobReqDeleteType{
			value: "roll_back_upgrade_active_standby",
		},
	}
}

func (c DeleteJobReqDeleteType) Value() string {
	return c.value
}

func (c DeleteJobReqDeleteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteJobReqDeleteType) UnmarshalJSON(b []byte) error {
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
