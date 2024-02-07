package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserDisclaimerRequest Request Object
type CreateUserDisclaimerRequest struct {
}

func (o CreateUserDisclaimerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserDisclaimerRequest struct{}"
	}

	return strings.Join([]string{"CreateUserDisclaimerRequest", string(data)}, " ")
}
