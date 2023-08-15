package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ApiParaForAuthToInstance struct {

	// api编号
	ApiId *string `json:"api_id,omitempty"`

	// 集群编号
	InstanceId *string `json:"instance_id,omitempty"`

	// app编号
	AppId *string `json:"app_id,omitempty"`

	// 申请类型
	ApplyType *ApiParaForAuthToInstanceApplyType `json:"apply_type,omitempty"`

	// 截止时间
	Time *string `json:"time,omitempty"`
}

func (o ApiParaForAuthToInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiParaForAuthToInstance struct{}"
	}

	return strings.Join([]string{"ApiParaForAuthToInstance", string(data)}, " ")
}

type ApiParaForAuthToInstanceApplyType struct {
	value string
}

type ApiParaForAuthToInstanceApplyTypeEnum struct {
	APPLY_TYPE_AUTHORIZE            ApiParaForAuthToInstanceApplyType
	APPLY_TYPE_API_CANCEL_AUTHORIZE ApiParaForAuthToInstanceApplyType
	APPLY_TYPE_APP_CANCEL_AUTHORIZE ApiParaForAuthToInstanceApplyType
	APPLY_TYPE_APPLY                ApiParaForAuthToInstanceApplyType
	APPLY_TYPE_RENEW                ApiParaForAuthToInstanceApplyType
}

func GetApiParaForAuthToInstanceApplyTypeEnum() ApiParaForAuthToInstanceApplyTypeEnum {
	return ApiParaForAuthToInstanceApplyTypeEnum{
		APPLY_TYPE_AUTHORIZE: ApiParaForAuthToInstanceApplyType{
			value: "APPLY_TYPE_AUTHORIZE",
		},
		APPLY_TYPE_API_CANCEL_AUTHORIZE: ApiParaForAuthToInstanceApplyType{
			value: "APPLY_TYPE_API_CANCEL_AUTHORIZE",
		},
		APPLY_TYPE_APP_CANCEL_AUTHORIZE: ApiParaForAuthToInstanceApplyType{
			value: "APPLY_TYPE_APP_CANCEL_AUTHORIZE",
		},
		APPLY_TYPE_APPLY: ApiParaForAuthToInstanceApplyType{
			value: "APPLY_TYPE_APPLY",
		},
		APPLY_TYPE_RENEW: ApiParaForAuthToInstanceApplyType{
			value: "APPLY_TYPE_RENEW",
		},
	}
}

func (c ApiParaForAuthToInstanceApplyType) Value() string {
	return c.value
}

func (c ApiParaForAuthToInstanceApplyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiParaForAuthToInstanceApplyType) UnmarshalJSON(b []byte) error {
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
