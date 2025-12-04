package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateValueListResponse Response Object
type UpdateValueListResponse struct {

	// 引用表id
	Id *string `json:"id,omitempty"`

	// 引用表名称
	Name *string `json:"name,omitempty"`

	// **参数解释：** 引用表类型 **约束限制：** 不涉及 **取值范围：**  - url  - params  - ip  - cookie  - referer  - user-agent  - header  - response_code  - response_header  - response_body  **默认取值：** 不涉及
	Type *UpdateValueListResponseType `json:"type,omitempty"`

	// 引用表描述
	Description *string `json:"description,omitempty"`

	// 引用表的值
	Values *[]string `json:"values,omitempty"`

	// 引用表来源：  - 1:表示来源于用户手动创建  - 2:表示来源于智能访问控制自动创建
	Producer       *int32 `json:"producer,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateValueListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateValueListResponse struct{}"
	}

	return strings.Join([]string{"UpdateValueListResponse", string(data)}, " ")
}

type UpdateValueListResponseType struct {
	value string
}

type UpdateValueListResponseTypeEnum struct {
	URL             UpdateValueListResponseType
	PARAMS          UpdateValueListResponseType
	IP              UpdateValueListResponseType
	COOKIE          UpdateValueListResponseType
	REFERER         UpdateValueListResponseType
	USER_AGENT      UpdateValueListResponseType
	HEADER          UpdateValueListResponseType
	RESPONSE_CODE   UpdateValueListResponseType
	RESPONSE_HEADER UpdateValueListResponseType
	RESPONSE_BODY   UpdateValueListResponseType
}

func GetUpdateValueListResponseTypeEnum() UpdateValueListResponseTypeEnum {
	return UpdateValueListResponseTypeEnum{
		URL: UpdateValueListResponseType{
			value: "url",
		},
		PARAMS: UpdateValueListResponseType{
			value: "params",
		},
		IP: UpdateValueListResponseType{
			value: "ip",
		},
		COOKIE: UpdateValueListResponseType{
			value: "cookie",
		},
		REFERER: UpdateValueListResponseType{
			value: "referer",
		},
		USER_AGENT: UpdateValueListResponseType{
			value: "user-agent",
		},
		HEADER: UpdateValueListResponseType{
			value: "header",
		},
		RESPONSE_CODE: UpdateValueListResponseType{
			value: "response_code",
		},
		RESPONSE_HEADER: UpdateValueListResponseType{
			value: "response_header",
		},
		RESPONSE_BODY: UpdateValueListResponseType{
			value: "response_body",
		},
	}
}

func (c UpdateValueListResponseType) Value() string {
	return c.value
}

func (c UpdateValueListResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateValueListResponseType) UnmarshalJSON(b []byte) error {
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
