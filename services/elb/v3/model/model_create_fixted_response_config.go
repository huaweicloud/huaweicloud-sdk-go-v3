package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateFixtedResponseConfig 参数解释：固定返回页面的配置。  约束限制： - 当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。 - 当action为FIXED_RESPONSE时生效，且为必选字段。其他action不可指定，否则报错。 [- 共享型负载均衡器下的转发策略不支持该字段，传入会报错。 ](tag:hws,hws_hk,ocb,ctc,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt,hk_tm)  [不支持该字段，请勿使用。](tag:hcso_dt)  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
type CreateFixtedResponseConfig struct {

	// 参数解释：返回码。  取值范围：200-299，400-499，500-599。
	StatusCode string `json:"status_code"`

	// 参数解释：返回body的格式。  取值范围： - text/plain - text/css - text/html - application/javascript - application/json  默认取值：text/plain
	ContentType *CreateFixtedResponseConfigContentType `json:"content_type,omitempty"`

	// 参数解释：返回消息内容。
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
