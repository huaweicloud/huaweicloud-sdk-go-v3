package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckUserIdentityRequest struct {
	Body *CheckSubcustomerUserReq `json:"body,omitempty" xml:"body"`
}

func (o CheckUserIdentityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckUserIdentityRequest struct{}"
	}

	return strings.Join([]string{"CheckUserIdentityRequest", string(data)}, " ")
}
