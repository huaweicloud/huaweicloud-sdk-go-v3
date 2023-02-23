package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowResourceRelationsDetailRequest struct {

	// 资源ID
	ResourceId string `json:"resource_id"`

	// 资源关系的指向。
	Direction ShowResourceRelationsDetailRequestDirection `json:"direction"`

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`
}

func (o ShowResourceRelationsDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceRelationsDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceRelationsDetailRequest", string(data)}, " ")
}

type ShowResourceRelationsDetailRequestDirection struct {
	value string
}

type ShowResourceRelationsDetailRequestDirectionEnum struct {
	IN  ShowResourceRelationsDetailRequestDirection
	OUT ShowResourceRelationsDetailRequestDirection
}

func GetShowResourceRelationsDetailRequestDirectionEnum() ShowResourceRelationsDetailRequestDirectionEnum {
	return ShowResourceRelationsDetailRequestDirectionEnum{
		IN: ShowResourceRelationsDetailRequestDirection{
			value: "in",
		},
		OUT: ShowResourceRelationsDetailRequestDirection{
			value: "out",
		},
	}
}

func (c ShowResourceRelationsDetailRequestDirection) Value() string {
	return c.value
}

func (c ShowResourceRelationsDetailRequestDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceRelationsDetailRequestDirection) UnmarshalJSON(b []byte) error {
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
