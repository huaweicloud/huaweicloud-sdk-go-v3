package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserDisclaimerRequest Request Object
type DeleteUserDisclaimerRequest struct {
}

func (o DeleteUserDisclaimerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserDisclaimerRequest struct{}"
	}

	return strings.Join([]string{"DeleteUserDisclaimerRequest", string(data)}, " ")
}
