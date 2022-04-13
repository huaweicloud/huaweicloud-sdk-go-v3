package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAccountStatusRequest struct {
}

func (o ShowAccountStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccountStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowAccountStatusRequest", string(data)}, " ")
}
