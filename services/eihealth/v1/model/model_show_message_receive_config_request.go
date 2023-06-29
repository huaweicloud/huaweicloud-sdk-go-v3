package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMessageReceiveConfigRequest Request Object
type ShowMessageReceiveConfigRequest struct {
}

func (o ShowMessageReceiveConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessageReceiveConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowMessageReceiveConfigRequest", string(data)}, " ")
}
