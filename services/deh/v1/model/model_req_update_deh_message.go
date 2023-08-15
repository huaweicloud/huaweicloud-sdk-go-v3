package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReqUpdateDehMessage 更新专属主机属性消息体。
type ReqUpdateDehMessage struct {

	// 在创建云服务器时（未指定专属主机ID），是否允许云服务器自动分配在一台可用的专属主机上。  取值范围：“on”或“off”。
	AutoPlacement *ReqUpdateDehMessageAutoPlacement `json:"auto_placement,omitempty"`

	// 专属主机名称。
	Name *string `json:"name,omitempty"`
}

func (o ReqUpdateDehMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqUpdateDehMessage struct{}"
	}

	return strings.Join([]string{"ReqUpdateDehMessage", string(data)}, " ")
}

type ReqUpdateDehMessageAutoPlacement struct {
	value string
}

type ReqUpdateDehMessageAutoPlacementEnum struct {
	ON  ReqUpdateDehMessageAutoPlacement
	OFF ReqUpdateDehMessageAutoPlacement
}

func GetReqUpdateDehMessageAutoPlacementEnum() ReqUpdateDehMessageAutoPlacementEnum {
	return ReqUpdateDehMessageAutoPlacementEnum{
		ON: ReqUpdateDehMessageAutoPlacement{
			value: "on",
		},
		OFF: ReqUpdateDehMessageAutoPlacement{
			value: "off",
		},
	}
}

func (c ReqUpdateDehMessageAutoPlacement) Value() string {
	return c.value
}

func (c ReqUpdateDehMessageAutoPlacement) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqUpdateDehMessageAutoPlacement) UnmarshalJSON(b []byte) error {
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
