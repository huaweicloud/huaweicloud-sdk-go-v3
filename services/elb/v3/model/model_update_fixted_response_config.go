package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 固定返回页面的配置。当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。共享型负载均衡器下的转发策略不支持该字段，传入会报错。  [当action为FIXED_RESPONSE时生效，且为必选字段，其他action不可指定，否则报错。](tag:hws,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42) [不支持该字段，请勿使用。](tag:dt,dt_test)
type UpdateFixtedResponseConfig struct {

	// 返回码。支持200~299,400~499,500~599。
	StatusCode *string `json:"status_code,omitempty"`

	// 返回body的格式。  取值范围： - text/plain - text/css - text/html - application/javascript - application/json
	ContentType *UpdateFixtedResponseConfigContentType `json:"content_type,omitempty"`

	// 返回消息内容。
	MessageBody *string `json:"message_body,omitempty"`
}

func (o UpdateFixtedResponseConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFixtedResponseConfig struct{}"
	}

	return strings.Join([]string{"UpdateFixtedResponseConfig", string(data)}, " ")
}

type UpdateFixtedResponseConfigContentType struct {
	value string
}

type UpdateFixtedResponseConfigContentTypeEnum struct {
	TEXT_PLAIN             UpdateFixtedResponseConfigContentType
	TEXT_CSS               UpdateFixtedResponseConfigContentType
	TEXT_HTML              UpdateFixtedResponseConfigContentType
	APPLICATION_JAVASCRIPT UpdateFixtedResponseConfigContentType
	APPLICATION_JSON       UpdateFixtedResponseConfigContentType
}

func GetUpdateFixtedResponseConfigContentTypeEnum() UpdateFixtedResponseConfigContentTypeEnum {
	return UpdateFixtedResponseConfigContentTypeEnum{
		TEXT_PLAIN: UpdateFixtedResponseConfigContentType{
			value: "text/plain",
		},
		TEXT_CSS: UpdateFixtedResponseConfigContentType{
			value: "text/css",
		},
		TEXT_HTML: UpdateFixtedResponseConfigContentType{
			value: "text/html",
		},
		APPLICATION_JAVASCRIPT: UpdateFixtedResponseConfigContentType{
			value: "application/javascript",
		},
		APPLICATION_JSON: UpdateFixtedResponseConfigContentType{
			value: "application/json",
		},
	}
}

func (c UpdateFixtedResponseConfigContentType) Value() string {
	return c.value
}

func (c UpdateFixtedResponseConfigContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFixtedResponseConfigContentType) UnmarshalJSON(b []byte) error {
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
