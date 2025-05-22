package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyAccountInfoRequest Request Object
type ModifyAccountInfoRequest struct {
	Body *AccountConfigModifyRequest `json:"body,omitempty"`
}

func (o ModifyAccountInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyAccountInfoRequest struct{}"
	}

	return strings.Join([]string{"ModifyAccountInfoRequest", string(data)}, " ")
}
