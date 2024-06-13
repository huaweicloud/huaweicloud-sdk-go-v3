package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemediationStaticParameter 合规规则修正执行的静态参数。
type RemediationStaticParameter struct {

	// 参数名称。
	VarKey *string `json:"var_key,omitempty"`

	// 参数的值。
	VarValue *interface{} `json:"var_value,omitempty"`
}

func (o RemediationStaticParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemediationStaticParameter struct{}"
	}

	return strings.Join([]string{"RemediationStaticParameter", string(data)}, " ")
}
