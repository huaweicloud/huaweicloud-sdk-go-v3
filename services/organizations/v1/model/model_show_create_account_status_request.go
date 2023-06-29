package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCreateAccountStatusRequest Request Object
type ShowCreateAccountStatusRequest struct {

	// 指定唯一标识CreateAccount请求的ID值。
	CreateAccountStatusId string `json:"create_account_status_id"`
}

func (o ShowCreateAccountStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCreateAccountStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowCreateAccountStatusRequest", string(data)}, " ")
}
