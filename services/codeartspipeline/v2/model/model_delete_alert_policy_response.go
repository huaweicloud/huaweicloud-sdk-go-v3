package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlertPolicyResponse Response Object
type DeleteAlertPolicyResponse struct {

	// **参数解释**： 操作是否成功。 **取值范围**： - true：操作成功。 - false：操作失败。
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteAlertPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlertPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlertPolicyResponse", string(data)}, " ")
}
