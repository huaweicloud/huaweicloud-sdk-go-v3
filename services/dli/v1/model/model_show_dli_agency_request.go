package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDliAgencyRequest Request Object
type ShowDliAgencyRequest struct {
}

func (o ShowDliAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDliAgencyRequest struct{}"
	}

	return strings.Join([]string{"ShowDliAgencyRequest", string(data)}, " ")
}
