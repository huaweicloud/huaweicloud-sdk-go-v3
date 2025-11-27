package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PrivatePolicyBodyPrimitiveTypeHolder struct {

	// 策略内容。仅支持OPA开源引擎识别的，以Rego（https://www.openpolicyagent.org/docs/latest/policy-language/）语言编写的策略模板。  policy_body和policy_uri 必须有且只有一个存在
	PolicyBody *string `json:"policy_body,omitempty"`
}

func (o PrivatePolicyBodyPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivatePolicyBodyPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"PrivatePolicyBodyPrimitiveTypeHolder", string(data)}, " ")
}
