package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CorpBasicDto 企业的基本信息。
type CorpBasicDto struct {

	// 企业名称，格式必须满足^[^#%&'+;<>=\\\"'？?\\\\\\\\……/]*$。
	Name string `json:"name"`

	// 企业域名。
	Domain *string `json:"domain,omitempty"`

	// 手机号，必须加上国家码，例如中国大陆手机+86xxxxxxx，当填写手机号时， “country”参数必填,手机格式必须满足(^$|^[+]?[0-9]+$)。
	Phone *string `json:"phone,omitempty"`

	// [[手机号所属的国家](https://support.huaweicloud.com/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hws)[[手机号所属的国家](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0109.html#ZH-CN_TOPIC_0212714591__table19371178135314)](tag:hk) 。
	Country *string `json:"country,omitempty"`

	// 传真号码,格式必须满足(^$|^[+]?[0-9]+$)。
	Fax *string `json:"fax,omitempty"`

	// 邮箱地址,格式必须满足(^$|^[\\\\w-+]+(\\\\.[\\\\w-+]+)*@[\\\\w-]+(\\\\.[\\\\w-]+)*(\\\\.[\\\\w-]{1,})$)。
	Email *string `json:"email,omitempty"`

	// 地址。
	Address *string `json:"address,omitempty"`

	// 备注。
	Description *string `json:"description,omitempty"`

	// 企业归属的SP ID。仅在查询时返回。
	SpId *string `json:"spId,omitempty"`

	// 企业提示音语言设置,zh-CN或en-US。
	Language *CorpBasicDtoLanguage `json:"language,omitempty"`

	// 时区Id设置,例如北京东8区timeZoneId值为56,时区Id和时区的对应关系请参考: [[时区表](https://support.huaweicloud.com/api-meeting/meeting_21_0110.html)](tag:hws)[[时区表](https://support.huaweicloud.com/intl/zh-cn/api-meeting/meeting_21_0110.html)](tag:hk) 。
	TimeZoneId *string `json:"timeZoneId,omitempty"`
}

func (o CorpBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CorpBasicDto struct{}"
	}

	return strings.Join([]string{"CorpBasicDto", string(data)}, " ")
}

type CorpBasicDtoLanguage struct {
	value string
}

type CorpBasicDtoLanguageEnum struct {
	ZH_CN CorpBasicDtoLanguage
	EN_US CorpBasicDtoLanguage
}

func GetCorpBasicDtoLanguageEnum() CorpBasicDtoLanguageEnum {
	return CorpBasicDtoLanguageEnum{
		ZH_CN: CorpBasicDtoLanguage{
			value: "zh-CN",
		},
		EN_US: CorpBasicDtoLanguage{
			value: "en-US",
		},
	}
}

func (c CorpBasicDtoLanguage) Value() string {
	return c.value
}

func (c CorpBasicDtoLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CorpBasicDtoLanguage) UnmarshalJSON(b []byte) error {
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
