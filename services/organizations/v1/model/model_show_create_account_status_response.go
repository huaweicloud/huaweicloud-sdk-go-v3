package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCreateAccountStatusResponse Response Object
type ShowCreateAccountStatusResponse struct {
	CreateAccountStatus *CreateAccountStatusDto `json:"create_account_status,omitempty"`
	HttpStatusCode      int                     `json:"-"`
}

func (o ShowCreateAccountStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCreateAccountStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowCreateAccountStatusResponse", string(data)}, " ")
}
