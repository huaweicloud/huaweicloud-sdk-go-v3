package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SignV2AgreementRequestBody struct {

	// **参数解释**：租户协议列表。 **约束限制**：最多10个协议。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Agreements []interface{} `json:"agreements"`
}

func (o SignV2AgreementRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignV2AgreementRequestBody struct{}"
	}

	return strings.Join([]string{"SignV2AgreementRequestBody", string(data)}, " ")
}
