package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchActionDesktopsReq 批量操作桌面请求。
type BatchActionDesktopsReq struct {

	// 操作的桌面ID列表。
	DesktopIds []string `json:"desktop_ids"`

	// 操作类型。 -os-start 启动。 -reboot 重启。 -os-stop 关机。 -os-hibernate 休眠。
	OpType string `json:"op_type"`

	// SOFT：普通操作；HARD：强制操作。例如type为HARD，op_type为os-stop代表强制关机。
	Type *BatchActionDesktopsReqType `json:"type,omitempty"`
}

func (o BatchActionDesktopsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchActionDesktopsReq struct{}"
	}

	return strings.Join([]string{"BatchActionDesktopsReq", string(data)}, " ")
}

type BatchActionDesktopsReqType struct {
	value string
}

type BatchActionDesktopsReqTypeEnum struct {
	SOFT BatchActionDesktopsReqType
	HARD BatchActionDesktopsReqType
}

func GetBatchActionDesktopsReqTypeEnum() BatchActionDesktopsReqTypeEnum {
	return BatchActionDesktopsReqTypeEnum{
		SOFT: BatchActionDesktopsReqType{
			value: "SOFT",
		},
		HARD: BatchActionDesktopsReqType{
			value: "HARD",
		},
	}
}

func (c BatchActionDesktopsReqType) Value() string {
	return c.value
}

func (c BatchActionDesktopsReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchActionDesktopsReqType) UnmarshalJSON(b []byte) error {
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
