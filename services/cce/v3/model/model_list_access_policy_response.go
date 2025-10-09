package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessPolicyResponse Response Object
type ListAccessPolicyResponse struct {

	// **参数解释：** API类型。 **约束限制：** 该值不可修改。 **取值范围：** 不涉及 **默认取值：** List
	Kind *string `json:"kind,omitempty"`

	// **参数解释：** API版本。 **约束限制：** 该值不可修改。 **取值范围：** 不涉及 **默认取值：** v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	AccessPolicyList *[]AccessPolicyResp `json:"accessPolicyList,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ListAccessPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListAccessPolicyResponse", string(data)}, " ")
}
