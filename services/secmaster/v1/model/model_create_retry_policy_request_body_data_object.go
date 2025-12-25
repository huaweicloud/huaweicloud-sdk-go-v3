package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRetryPolicyRequestBodyDataObject 策略实体信息
type CreateRetryPolicyRequestBodyDataObject struct {

	// 重试策略ID
	RetryList *[]string `json:"retry_list,omitempty"`

	BlockAge *CreateRetryPolicyRequestBodyDataObjectBlockAge `json:"block_age"`

	// 策略对象
	BlockTarget string `json:"block_target"`

	// 与操作连接对应的策略列表
	DefensePolicyList []CreateRetryPolicyRequestBodyDataObjectDefensePolicyList `json:"defense_policy_list"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 标签
	Labels *string `json:"labels,omitempty"`

	// 类型,WHITE/BLOCK,WHITE代表加白(将ip等对象加入白名单),BLOCK代表阻断(将ip等对象加入黑名单)
	PolicyCategory CreateRetryPolicyRequestBodyDataObjectPolicyCategory `json:"policy_category"`

	PolicyType *CreateRetryPolicyRequestBodyDataObjectPolicyType `json:"policy_type"`

	// 区域ID
	RegionId string `json:"region_id"`

	// 出入方向
	PolicyDirection *string `json:"policy_direction,omitempty"`

	// 账号范围
	AccountScope *string `json:"account_scope,omitempty"`

	// 企业项目范围
	EpsScope *string `json:"eps_scope,omitempty"`

	// region范围
	RegionScope *string `json:"region_scope,omitempty"`
}

func (o CreateRetryPolicyRequestBodyDataObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetryPolicyRequestBodyDataObject struct{}"
	}

	return strings.Join([]string{"CreateRetryPolicyRequestBodyDataObject", string(data)}, " ")
}

type CreateRetryPolicyRequestBodyDataObjectPolicyCategory struct {
	value string
}

type CreateRetryPolicyRequestBodyDataObjectPolicyCategoryEnum struct {
	WHITE CreateRetryPolicyRequestBodyDataObjectPolicyCategory
	BLOCK CreateRetryPolicyRequestBodyDataObjectPolicyCategory
}

func GetCreateRetryPolicyRequestBodyDataObjectPolicyCategoryEnum() CreateRetryPolicyRequestBodyDataObjectPolicyCategoryEnum {
	return CreateRetryPolicyRequestBodyDataObjectPolicyCategoryEnum{
		WHITE: CreateRetryPolicyRequestBodyDataObjectPolicyCategory{
			value: "WHITE",
		},
		BLOCK: CreateRetryPolicyRequestBodyDataObjectPolicyCategory{
			value: "BLOCK",
		},
	}
}

func (c CreateRetryPolicyRequestBodyDataObjectPolicyCategory) Value() string {
	return c.value
}

func (c CreateRetryPolicyRequestBodyDataObjectPolicyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRetryPolicyRequestBodyDataObjectPolicyCategory) UnmarshalJSON(b []byte) error {
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
