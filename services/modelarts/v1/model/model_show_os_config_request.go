package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOsConfigRequest Request Object
type ShowOsConfigRequest struct {
}

func (o ShowOsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOsConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowOsConfigRequest", string(data)}, " ")
}
