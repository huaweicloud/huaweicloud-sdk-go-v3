package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 固定返回页面的配置。当监听器的高级转发策略功能（enhance_l7policy_enable）开启后才会生效，未开启传入该字段会报错。共享型负载均衡器下的转发策略不支持该字段，传入会报错。 [当action为FIXED_RESPONSE时生效，且为必选字段，其他action不可指定。](tag:hc,hws,hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42) [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)
type FixtedResponseConfig struct {
	// 返回码。支持200~299,400~499,500~599。

	StatusCode string `json:"status_code"`
	// 返回body的格式。  取值范围： - text/plain - text/css - text/html - application/javascript - application/json

	ContentType FixtedResponseConfigContentType `json:"content_type"`
	// 返回消息内容。

	MessageBody string `json:"message_body"`
}

func (o FixtedResponseConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FixtedResponseConfig struct{}"
	}

	return strings.Join([]string{"FixtedResponseConfig", string(data)}, " ")
}

type FixtedResponseConfigContentType struct {
	value string
}

type FixtedResponseConfigContentTypeEnum struct {
	TEXT_PLAIN             FixtedResponseConfigContentType
	TEXT_CSS               FixtedResponseConfigContentType
	TEXT_HTML              FixtedResponseConfigContentType
	APPLICATION_JAVASCRIPT FixtedResponseConfigContentType
	APPLICATION_JSON       FixtedResponseConfigContentType
}

func GetFixtedResponseConfigContentTypeEnum() FixtedResponseConfigContentTypeEnum {
	return FixtedResponseConfigContentTypeEnum{
		TEXT_PLAIN: FixtedResponseConfigContentType{
			value: "text/plain",
		},
		TEXT_CSS: FixtedResponseConfigContentType{
			value: "text/css",
		},
		TEXT_HTML: FixtedResponseConfigContentType{
			value: "text/html",
		},
		APPLICATION_JAVASCRIPT: FixtedResponseConfigContentType{
			value: "application/javascript",
		},
		APPLICATION_JSON: FixtedResponseConfigContentType{
			value: "application/json",
		},
	}
}

func (c FixtedResponseConfigContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FixtedResponseConfigContentType) UnmarshalJSON(b []byte) error {
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
