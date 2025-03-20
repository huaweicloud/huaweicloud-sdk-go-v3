package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowLatestPublishedAgreementRequest Request Object
type ShowLatestPublishedAgreementRequest struct {

	// 协议类型，0 新建工单-普通用户，1 新建工单-ISV用户，2 授权
	AgreementType *ShowLatestPublishedAgreementRequestAgreementType `json:"agreement_type,omitempty"`

	// 对接站点信息。  0（中国站） 1（国际站），不填的话默认为0。
	XSite *int32 `json:"X-Site,omitempty"`

	// 语言环境，值为通用的语言描述字符串，比如zh-cn等，默认为zh-cn。  会根据语言环境对应展示一些国际化的信息，比如工单类型名称等。
	XLanguage *string `json:"X-Language,omitempty"`

	// 环境时区，值为通用的时区描述字符串，比如GMT+8等，默认为GMT+8。  涉及时间的数据会根据环境时区处理。
	XTimeZone *string `json:"X-Time-Zone,omitempty"`
}

func (o ShowLatestPublishedAgreementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestPublishedAgreementRequest struct{}"
	}

	return strings.Join([]string{"ShowLatestPublishedAgreementRequest", string(data)}, " ")
}

type ShowLatestPublishedAgreementRequestAgreementType struct {
	value string
}

type ShowLatestPublishedAgreementRequestAgreementTypeEnum struct {
	E_0 ShowLatestPublishedAgreementRequestAgreementType
	E_1 ShowLatestPublishedAgreementRequestAgreementType
	E_2 ShowLatestPublishedAgreementRequestAgreementType
}

func GetShowLatestPublishedAgreementRequestAgreementTypeEnum() ShowLatestPublishedAgreementRequestAgreementTypeEnum {
	return ShowLatestPublishedAgreementRequestAgreementTypeEnum{
		E_0: ShowLatestPublishedAgreementRequestAgreementType{
			value: "0",
		},
		E_1: ShowLatestPublishedAgreementRequestAgreementType{
			value: "1",
		},
		E_2: ShowLatestPublishedAgreementRequestAgreementType{
			value: "2",
		},
	}
}

func (c ShowLatestPublishedAgreementRequestAgreementType) Value() string {
	return c.value
}

func (c ShowLatestPublishedAgreementRequestAgreementType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLatestPublishedAgreementRequestAgreementType) UnmarshalJSON(b []byte) error {
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
