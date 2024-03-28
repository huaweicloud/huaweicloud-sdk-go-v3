package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GcbAdminState struct {

	// 功能说明: 全域互联带宽状态。 取值范围：     NORMAL-正常     FREEZED-冻结状态
	AdminState *GcbAdminStateAdminState `json:"admin_state,omitempty"`
}

func (o GcbAdminState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbAdminState struct{}"
	}

	return strings.Join([]string{"GcbAdminState", string(data)}, " ")
}

type GcbAdminStateAdminState struct {
	value string
}

type GcbAdminStateAdminStateEnum struct {
	NORMAL  GcbAdminStateAdminState
	FREEZED GcbAdminStateAdminState
}

func GetGcbAdminStateAdminStateEnum() GcbAdminStateAdminStateEnum {
	return GcbAdminStateAdminStateEnum{
		NORMAL: GcbAdminStateAdminState{
			value: "NORMAL",
		},
		FREEZED: GcbAdminStateAdminState{
			value: "FREEZED",
		},
	}
}

func (c GcbAdminStateAdminState) Value() string {
	return c.value
}

func (c GcbAdminStateAdminState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GcbAdminStateAdminState) UnmarshalJSON(b []byte) error {
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
