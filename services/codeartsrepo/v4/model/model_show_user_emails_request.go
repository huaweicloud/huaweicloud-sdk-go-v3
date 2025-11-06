package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserEmailsRequest Request Object
type ShowUserEmailsRequest struct {
}

func (o ShowUserEmailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserEmailsRequest struct{}"
	}

	return strings.Join([]string{"ShowUserEmailsRequest", string(data)}, " ")
}
