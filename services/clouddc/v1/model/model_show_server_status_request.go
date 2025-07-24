package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerStatusRequest Request Object
type ShowServerStatusRequest struct {
}

func (o ShowServerStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowServerStatusRequest", string(data)}, " ")
}
