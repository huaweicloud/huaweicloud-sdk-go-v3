package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateFixtedResponseConfig struct {

	// **参数解释**：返回码。  **约束限制**：不涉及  **取值范围**：200-299，400-499，500-599。  **默认取值**：不涉及
	StatusCode string `json:"status_code"`

	// **参数解释**：返回body的格式。  **约束限制**：不涉及  **取值范围**： - text/plain - text/css - text/html - application/javascript - application/json  **默认取值**：不涉及
	ContentType *CreateFixtedResponseConfigContentType `json:"content_type,omitempty"`

	// **参数解释**：返回消息内容。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	MessageBody *string `json:"message_body,omitempty"`

	InsertHeadersConfig *CreateInsertHeadersConfig `json:"insert_headers_config,omitempty"`

	RemoveHeadersConfig *CreateRemoveHeadersConfig `json:"remove_headers_config,omitempty"`

	TrafficLimitConfig *CreateTrafficLimitConfig `json:"traffic_limit_config,omitempty"`
}

func (o CreateFixtedResponseConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFixtedResponseConfig struct{}"
	}

	return strings.Join([]string{"CreateFixtedResponseConfig", string(data)}, " ")
}

type CreateFixtedResponseConfigContentType struct {
	value string
}

type CreateFixtedResponseConfigContentTypeEnum struct {
	TEXT_PLAIN             CreateFixtedResponseConfigContentType
	TEXT_CSS               CreateFixtedResponseConfigContentType
	TEXT_HTML              CreateFixtedResponseConfigContentType
	APPLICATION_JAVASCRIPT CreateFixtedResponseConfigContentType
	APPLICATION_JSON       CreateFixtedResponseConfigContentType
}

func GetCreateFixtedResponseConfigContentTypeEnum() CreateFixtedResponseConfigContentTypeEnum {
	return CreateFixtedResponseConfigContentTypeEnum{
		TEXT_PLAIN: CreateFixtedResponseConfigContentType{
			value: "text/plain",
		},
		TEXT_CSS: CreateFixtedResponseConfigContentType{
			value: "text/css",
		},
		TEXT_HTML: CreateFixtedResponseConfigContentType{
			value: "text/html",
		},
		APPLICATION_JAVASCRIPT: CreateFixtedResponseConfigContentType{
			value: "application/javascript",
		},
		APPLICATION_JSON: CreateFixtedResponseConfigContentType{
			value: "application/json",
		},
	}
}

func (c CreateFixtedResponseConfigContentType) Value() string {
	return c.value
}

func (c CreateFixtedResponseConfigContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFixtedResponseConfigContentType) UnmarshalJSON(b []byte) error {
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
