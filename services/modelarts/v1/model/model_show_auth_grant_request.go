package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuthGrantRequest Request Object
type ShowAuthGrantRequest struct {
}

func (o ShowAuthGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuthGrantRequest struct{}"
	}

	return strings.Join([]string{"ShowAuthGrantRequest", string(data)}, " ")
}
