package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSourceIpRequest Request Object
type ShowSourceIpRequest struct {
}

func (o ShowSourceIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSourceIpRequest struct{}"
	}

	return strings.Join([]string{"ShowSourceIpRequest", string(data)}, " ")
}
