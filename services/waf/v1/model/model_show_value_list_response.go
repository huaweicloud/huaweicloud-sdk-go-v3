package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowValueListResponse Response Object
type ShowValueListResponse struct {

	// 引用表id
	Id *string `json:"id,omitempty"`

	// 引用表名称
	Name *string `json:"name,omitempty"`

	// **参数解释：** 引用表类型 **约束限制：** 不涉及 **取值范围：**  - url  - params  - ip  - cookie  - referer  - user-agent  - header  - response_code  - response_header  - response_body  **默认取值：** 不涉及
	Type *ShowValueListResponseType `json:"type,omitempty"`

	// 引用表描述
	Description *string `json:"description,omitempty"`

	// 引用表的值
	Values *[]string `json:"values,omitempty"`

	// 引用表来源：  - 1:表示来源于用户手动创建  - 2:表示来源于智能访问控制自动创建
	Producer *int32 `json:"producer,omitempty"`

	// 创建规则的时间，格式为13位毫秒时间戳
	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowValueListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowValueListResponse struct{}"
	}

	return strings.Join([]string{"ShowValueListResponse", string(data)}, " ")
}

type ShowValueListResponseType struct {
	value string
}

type ShowValueListResponseTypeEnum struct {
	URL             ShowValueListResponseType
	PARAMS          ShowValueListResponseType
	IP              ShowValueListResponseType
	COOKIE          ShowValueListResponseType
	REFERER         ShowValueListResponseType
	USER_AGENT      ShowValueListResponseType
	HEADER          ShowValueListResponseType
	RESPONSE_CODE   ShowValueListResponseType
	RESPONSE_HEADER ShowValueListResponseType
	RESPONSE_BODY   ShowValueListResponseType
}

func GetShowValueListResponseTypeEnum() ShowValueListResponseTypeEnum {
	return ShowValueListResponseTypeEnum{
		URL: ShowValueListResponseType{
			value: "url",
		},
		PARAMS: ShowValueListResponseType{
			value: "params",
		},
		IP: ShowValueListResponseType{
			value: "ip",
		},
		COOKIE: ShowValueListResponseType{
			value: "cookie",
		},
		REFERER: ShowValueListResponseType{
			value: "referer",
		},
		USER_AGENT: ShowValueListResponseType{
			value: "user-agent",
		},
		HEADER: ShowValueListResponseType{
			value: "header",
		},
		RESPONSE_CODE: ShowValueListResponseType{
			value: "response_code",
		},
		RESPONSE_HEADER: ShowValueListResponseType{
			value: "response_header",
		},
		RESPONSE_BODY: ShowValueListResponseType{
			value: "response_body",
		},
	}
}

func (c ShowValueListResponseType) Value() string {
	return c.value
}

func (c ShowValueListResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowValueListResponseType) UnmarshalJSON(b []byte) error {
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
