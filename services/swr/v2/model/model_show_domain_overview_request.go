package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDomainOverviewRequest Request Object
type ShowDomainOverviewRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ShowDomainOverviewRequestContentType `json:"Content-Type"`
}

func (o ShowDomainOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainOverviewRequest", string(data)}, " ")
}

type ShowDomainOverviewRequestContentType struct {
	value string
}

type ShowDomainOverviewRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowDomainOverviewRequestContentType
	APPLICATION_JSON             ShowDomainOverviewRequestContentType
}

func GetShowDomainOverviewRequestContentTypeEnum() ShowDomainOverviewRequestContentTypeEnum {
	return ShowDomainOverviewRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowDomainOverviewRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowDomainOverviewRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowDomainOverviewRequestContentType) Value() string {
	return c.value
}

func (c ShowDomainOverviewRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDomainOverviewRequestContentType) UnmarshalJSON(b []byte) error {
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
