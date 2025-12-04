package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateValueListResponse Response Object
type CreateValueListResponse struct {

	// 引用表id
	Id *string `json:"id,omitempty"`

	// 引用表名称
	Name *string `json:"name,omitempty"`

	// **参数解释：** 引用表类型 **约束限制：** 不涉及 **取值范围：**  - url  - params  - ip  - cookie  - referer  - user-agent  - header  - response_code  - response_header  - response_body  **默认取值：** 不涉及
	Type *CreateValueListResponseType `json:"type,omitempty"`

	// 引用表描述
	Description *string `json:"description,omitempty"`

	// 引用表时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 引用表的值
	Values *[]string `json:"values,omitempty"`

	// 引用表来源：  - 1:表示来源于用户手动创建  - 2:表示来源于智能访问控制自动创建
	Producer       *int32 `json:"producer,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateValueListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateValueListResponse struct{}"
	}

	return strings.Join([]string{"CreateValueListResponse", string(data)}, " ")
}

type CreateValueListResponseType struct {
	value string
}

type CreateValueListResponseTypeEnum struct {
	URL             CreateValueListResponseType
	PARAMS          CreateValueListResponseType
	IP              CreateValueListResponseType
	COOKIE          CreateValueListResponseType
	REFERER         CreateValueListResponseType
	USER_AGENT      CreateValueListResponseType
	HEADER          CreateValueListResponseType
	RESPONSE_CODE   CreateValueListResponseType
	RESPONSE_HEADER CreateValueListResponseType
	RESPONSE_BODY   CreateValueListResponseType
}

func GetCreateValueListResponseTypeEnum() CreateValueListResponseTypeEnum {
	return CreateValueListResponseTypeEnum{
		URL: CreateValueListResponseType{
			value: "url",
		},
		PARAMS: CreateValueListResponseType{
			value: "params",
		},
		IP: CreateValueListResponseType{
			value: "ip",
		},
		COOKIE: CreateValueListResponseType{
			value: "cookie",
		},
		REFERER: CreateValueListResponseType{
			value: "referer",
		},
		USER_AGENT: CreateValueListResponseType{
			value: "user-agent",
		},
		HEADER: CreateValueListResponseType{
			value: "header",
		},
		RESPONSE_CODE: CreateValueListResponseType{
			value: "response_code",
		},
		RESPONSE_HEADER: CreateValueListResponseType{
			value: "response_header",
		},
		RESPONSE_BODY: CreateValueListResponseType{
			value: "response_body",
		},
	}
}

func (c CreateValueListResponseType) Value() string {
	return c.value
}

func (c CreateValueListResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateValueListResponseType) UnmarshalJSON(b []byte) error {
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
