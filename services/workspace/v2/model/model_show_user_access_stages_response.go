package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowUserAccessStagesResponse Response Object
type ShowUserAccessStagesResponse struct {

	// 用户名。
	Username *string `json:"username,omitempty"`

	// 接入阶段 | APP、SHARE_APP - 应用 DESKTOP、VDI - 桌面。
	BizType *ShowUserAccessStagesResponseBizType `json:"biz_type,omitempty"`

	// 接入各阶段详情。
	Stages         *[]UserAccessStage `json:"stages,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowUserAccessStagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserAccessStagesResponse struct{}"
	}

	return strings.Join([]string{"ShowUserAccessStagesResponse", string(data)}, " ")
}

type ShowUserAccessStagesResponseBizType struct {
	value string
}

type ShowUserAccessStagesResponseBizTypeEnum struct {
	APP       ShowUserAccessStagesResponseBizType
	DESKTOP   ShowUserAccessStagesResponseBizType
	SHARE_APP ShowUserAccessStagesResponseBizType
	VDI       ShowUserAccessStagesResponseBizType
}

func GetShowUserAccessStagesResponseBizTypeEnum() ShowUserAccessStagesResponseBizTypeEnum {
	return ShowUserAccessStagesResponseBizTypeEnum{
		APP: ShowUserAccessStagesResponseBizType{
			value: "APP",
		},
		DESKTOP: ShowUserAccessStagesResponseBizType{
			value: "DESKTOP",
		},
		SHARE_APP: ShowUserAccessStagesResponseBizType{
			value: "SHARE_APP",
		},
		VDI: ShowUserAccessStagesResponseBizType{
			value: "VDI",
		},
	}
}

func (c ShowUserAccessStagesResponseBizType) Value() string {
	return c.value
}

func (c ShowUserAccessStagesResponseBizType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUserAccessStagesResponseBizType) UnmarshalJSON(b []byte) error {
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
