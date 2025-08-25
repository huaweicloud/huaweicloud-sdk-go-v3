package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableAzRequest Request Object
type ShowAvailableAzRequest struct {
}

func (o ShowAvailableAzRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableAzRequest struct{}"
	}

	return strings.Join([]string{"ShowAvailableAzRequest", string(data)}, " ")
}
