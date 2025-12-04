package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchCreatePrivacyRuleRequestBody struct {

	// 隐私屏蔽规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/_*,以\"*\"（星号）结尾代表路径前缀
	Url string `json:"url"`

	// **参数解释：** 屏蔽字段 **约束限制：** 不涉及 **取值范围：**  - params: 请求参数  - cookie: 根据Cookie区分的Web访问者  - header: 自定义HTTP首部  - form: 表单参数  **默认取值：** 不涉及
	Category BatchCreatePrivacyRuleRequestBodyCategory `json:"category"`

	// 屏蔽字段名，根据“屏蔽字段”设置字段名，被屏蔽的字段将不会出现在日志中。屏蔽字段名的长度不能超过2048字节，且只能由数字、字母、下划线和中划线组成。
	Index string `json:"index"`

	// 规则描述，可选参数，设置该规则的备注信息。
	Description *string `json:"description,omitempty"`

	// 添加规则的策略id列表。策略id从\"查询防护策略列表\"(ListPolicy)接口获取，多个策略之间用“,”隔开
	PolicyIds []string `json:"policy_ids"`
}

func (o BatchCreatePrivacyRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePrivacyRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreatePrivacyRuleRequestBody", string(data)}, " ")
}

type BatchCreatePrivacyRuleRequestBodyCategory struct {
	value string
}

type BatchCreatePrivacyRuleRequestBodyCategoryEnum struct {
	PARAMS BatchCreatePrivacyRuleRequestBodyCategory
	COOKIE BatchCreatePrivacyRuleRequestBodyCategory
	HEADER BatchCreatePrivacyRuleRequestBodyCategory
	FORM   BatchCreatePrivacyRuleRequestBodyCategory
}

func GetBatchCreatePrivacyRuleRequestBodyCategoryEnum() BatchCreatePrivacyRuleRequestBodyCategoryEnum {
	return BatchCreatePrivacyRuleRequestBodyCategoryEnum{
		PARAMS: BatchCreatePrivacyRuleRequestBodyCategory{
			value: "params",
		},
		COOKIE: BatchCreatePrivacyRuleRequestBodyCategory{
			value: "cookie",
		},
		HEADER: BatchCreatePrivacyRuleRequestBodyCategory{
			value: "header",
		},
		FORM: BatchCreatePrivacyRuleRequestBodyCategory{
			value: "form",
		},
	}
}

func (c BatchCreatePrivacyRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c BatchCreatePrivacyRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreatePrivacyRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
