package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAccountStatusResponse struct {
	Result *AccountStatus `json:"result,omitempty"`
	// 状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAccountStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccountStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAccountStatusResponse", string(data)}, " ")
}
