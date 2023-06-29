package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TmsUpdatePublicReq struct {

	// 操作标识：create（创建），delete（删除）
	Action TmsUpdatePublicReqAction `json:"action"`

	// 标签列表。  一个实例默认最多支持创建20个标签。
	Tags []TmsKeyValue `json:"tags"`
}

func (o TmsUpdatePublicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsUpdatePublicReq struct{}"
	}

	return strings.Join([]string{"TmsUpdatePublicReq", string(data)}, " ")
}

type TmsUpdatePublicReqAction struct {
	value string
}

type TmsUpdatePublicReqActionEnum struct {
	CREATE TmsUpdatePublicReqAction
	DELETE TmsUpdatePublicReqAction
}

func GetTmsUpdatePublicReqActionEnum() TmsUpdatePublicReqActionEnum {
	return TmsUpdatePublicReqActionEnum{
		CREATE: TmsUpdatePublicReqAction{
			value: "create",
		},
		DELETE: TmsUpdatePublicReqAction{
			value: "delete",
		},
	}
}

func (c TmsUpdatePublicReqAction) Value() string {
	return c.value
}

func (c TmsUpdatePublicReqAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TmsUpdatePublicReqAction) UnmarshalJSON(b []byte) error {
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
