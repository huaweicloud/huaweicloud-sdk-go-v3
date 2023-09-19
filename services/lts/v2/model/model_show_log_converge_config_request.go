package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogConvergeConfigRequest Request Object
type ShowLogConvergeConfigRequest struct {

	// 成员帐户ID
	MemberAccountId string `json:"member_account_id"`
}

func (o ShowLogConvergeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogConvergeConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowLogConvergeConfigRequest", string(data)}, " ")
}
