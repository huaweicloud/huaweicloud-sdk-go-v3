package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExpansionParametersRequestBody struct {

	// **参数解释**: 参数名称与参数值的键值对。 **取值范围**: 不涉及。
	Params map[string]string `json:"params"`
}

func (o ExpansionParametersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpansionParametersRequestBody struct{}"
	}

	return strings.Join([]string{"ExpansionParametersRequestBody", string(data)}, " ")
}
