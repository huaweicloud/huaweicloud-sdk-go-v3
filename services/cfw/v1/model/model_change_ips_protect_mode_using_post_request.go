package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeIpsProtectModeUsingPostRequest struct {
	Body *IpsProtectDto `json:"body,omitempty"`
}

func (o ChangeIpsProtectModeUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIpsProtectModeUsingPostRequest struct{}"
	}

	return strings.Join([]string{"ChangeIpsProtectModeUsingPostRequest", string(data)}, " ")
}
