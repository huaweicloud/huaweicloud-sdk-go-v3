package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowResourceRelationsRequest struct {
	// 资源ID

	ResourceId string `json:"resource_id"`
	// 资源关系的指向

	Direction ShowResourceRelationsRequestDirection `json:"direction"`
	// 最大的返回数量

	Limit *int32 `json:"limit,omitempty"`
	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页

	Marker *string `json:"marker,omitempty"`
}

func (o ShowResourceRelationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceRelationsRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceRelationsRequest", string(data)}, " ")
}

type ShowResourceRelationsRequestDirection struct {
	value string
}

type ShowResourceRelationsRequestDirectionEnum struct {
	IN  ShowResourceRelationsRequestDirection
	OUT ShowResourceRelationsRequestDirection
}

func GetShowResourceRelationsRequestDirectionEnum() ShowResourceRelationsRequestDirectionEnum {
	return ShowResourceRelationsRequestDirectionEnum{
		IN: ShowResourceRelationsRequestDirection{
			value: "in",
		},
		OUT: ShowResourceRelationsRequestDirection{
			value: "out",
		},
	}
}

func (c ShowResourceRelationsRequestDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceRelationsRequestDirection) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
