package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowElbIpGroupResponse Response Object
type ShowElbIpGroupResponse struct {

	// 类型选项，取值： - whiteList：白名单，只允许指定ip或网段访问。 - blackList：黑名单，不允许指定ip或网段访问。
	Type *ShowElbIpGroupResponseType `json:"type,omitempty"`

	// IP地址组中包含的IP或网段列表。
	IpGroups *[]SwitchIpGroupRequestBodyIpGroups `json:"ip_groups,omitempty"`

	// true：开启，false：关闭。
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowElbIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowElbIpGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowElbIpGroupResponse", string(data)}, " ")
}

type ShowElbIpGroupResponseType struct {
	value string
}

type ShowElbIpGroupResponseTypeEnum struct {
	WHITE_LIST ShowElbIpGroupResponseType
	BLACK_LIST ShowElbIpGroupResponseType
}

func GetShowElbIpGroupResponseTypeEnum() ShowElbIpGroupResponseTypeEnum {
	return ShowElbIpGroupResponseTypeEnum{
		WHITE_LIST: ShowElbIpGroupResponseType{
			value: "whiteList",
		},
		BLACK_LIST: ShowElbIpGroupResponseType{
			value: "blackList",
		},
	}
}

func (c ShowElbIpGroupResponseType) Value() string {
	return c.value
}

func (c ShowElbIpGroupResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowElbIpGroupResponseType) UnmarshalJSON(b []byte) error {
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
