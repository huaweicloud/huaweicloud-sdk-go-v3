package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ThirdTemplateRequestBody struct {

	// 必须以字母或数字开头，只能由字母、数字、下划线和中划线组成，长度小于等于64个字符。
	Category string `json:"category"`

	// Inputs参数
	Inputs *[]Input `json:"inputs,omitempty"`

	// 根据DWR自定义的函数模板创建属于用户的function，并指定该参数设置的参数值。
	DynamicSourceDefinition map[string]interface{} `json:"dynamic_source_definition,omitempty"`

	// 算子执行时需要的权限信息。
	NeedPolicy *[]Policy `json:"need_policy,omitempty"`

	// 算子提供方。 英文：必须以字母或数字开头，只能由字母、数字、下划线和中划线组成，长度小于等于32个字符。 中文：只能由中文、字母、数字、下划线和中划线组成，长度小于等于32个字符。
	FuncProvider map[string]string `json:"func_provider"`

	// 算子名称。 英文：必须以字母或数字开头，只能由字母、数字、下划线和中划线组成，长度小于等于50个字符。 中文：只能由中文、字母、数字、下划线和中划线组成，长度小于等于50个字符。
	FuncName map[string]string `json:"func_name"`

	// 算子描述。 英文：长度最小为0，最长为256，可以是数字、大小写字母以及英文的逗号，句号，冒号，中划线,，下划线，空格。 中文：长度最小为0，最长为256，可以是中文、数字、大小写字母以及中英文的逗号，句号，冒号，中划线，下划线，空格。
	FuncDescription map[string]string `json:"func_description"`

	// 云市场链接。 需要包含中文， 长度最小为0，最长为512，须遵守http协议中定义的规则。
	FuncLink map[string]string `json:"func_link,omitempty"`

	// serverless算子应用程序urn
	AppUrn *string `json:"app_urn,omitempty"`

	// Serverless计费单价
	BillValue *float64 `json:"bill_value,omitempty"`

	// serverless所需要委托名
	Agency *string `json:"agency,omitempty"`

	// 算子状态。初始创建时为init_created。申请提交时传入submit_approve。 算子状态。申请提交时传入submit_approve。
	RegisterStatus *ThirdTemplateRequestBodyRegisterStatus `json:"register_status,omitempty"`
}

func (o ThirdTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThirdTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"ThirdTemplateRequestBody", string(data)}, " ")
}

type ThirdTemplateRequestBodyRegisterStatus struct {
	value string
}

type ThirdTemplateRequestBodyRegisterStatusEnum struct {
	SUBMIT_APPROVE ThirdTemplateRequestBodyRegisterStatus
	INIT_CREATED   ThirdTemplateRequestBodyRegisterStatus
}

func GetThirdTemplateRequestBodyRegisterStatusEnum() ThirdTemplateRequestBodyRegisterStatusEnum {
	return ThirdTemplateRequestBodyRegisterStatusEnum{
		SUBMIT_APPROVE: ThirdTemplateRequestBodyRegisterStatus{
			value: "submit_approve",
		},
		INIT_CREATED: ThirdTemplateRequestBodyRegisterStatus{
			value: "init_created",
		},
	}
}

func (c ThirdTemplateRequestBodyRegisterStatus) Value() string {
	return c.value
}

func (c ThirdTemplateRequestBodyRegisterStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThirdTemplateRequestBodyRegisterStatus) UnmarshalJSON(b []byte) error {
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
