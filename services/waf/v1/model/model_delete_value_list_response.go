package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteValueListResponse Response Object
type DeleteValueListResponse struct {

	// 引用表id
	Id *string `json:"id,omitempty"`

	// 引用表名称
	Name *string `json:"name,omitempty"`

	// **参数解释：** 引用表类型 **约束限制：** 不涉及 **取值范围：**  - url  - params  - ip  - cookie  - referer  - user-agent  - header  - response_code  - response_header  - response_body  **默认取值：** 不涉及
	Type *DeleteValueListResponseType `json:"type,omitempty"`

	// 删除引用表的时间，时间为13位毫秒时间戳
	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteValueListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteValueListResponse struct{}"
	}

	return strings.Join([]string{"DeleteValueListResponse", string(data)}, " ")
}

type DeleteValueListResponseType struct {
	value string
}

type DeleteValueListResponseTypeEnum struct {
	URL             DeleteValueListResponseType
	PARAMS          DeleteValueListResponseType
	IP              DeleteValueListResponseType
	COOKIE          DeleteValueListResponseType
	REFERER         DeleteValueListResponseType
	USER_AGENT      DeleteValueListResponseType
	HEADER          DeleteValueListResponseType
	RESPONSE_CODE   DeleteValueListResponseType
	RESPONSE_HEADER DeleteValueListResponseType
	RESPONSE_BODY   DeleteValueListResponseType
}

func GetDeleteValueListResponseTypeEnum() DeleteValueListResponseTypeEnum {
	return DeleteValueListResponseTypeEnum{
		URL: DeleteValueListResponseType{
			value: "url",
		},
		PARAMS: DeleteValueListResponseType{
			value: "params",
		},
		IP: DeleteValueListResponseType{
			value: "ip",
		},
		COOKIE: DeleteValueListResponseType{
			value: "cookie",
		},
		REFERER: DeleteValueListResponseType{
			value: "referer",
		},
		USER_AGENT: DeleteValueListResponseType{
			value: "user-agent",
		},
		HEADER: DeleteValueListResponseType{
			value: "header",
		},
		RESPONSE_CODE: DeleteValueListResponseType{
			value: "response_code",
		},
		RESPONSE_HEADER: DeleteValueListResponseType{
			value: "response_header",
		},
		RESPONSE_BODY: DeleteValueListResponseType{
			value: "response_body",
		},
	}
}

func (c DeleteValueListResponseType) Value() string {
	return c.value
}

func (c DeleteValueListResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteValueListResponseType) UnmarshalJSON(b []byte) error {
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
