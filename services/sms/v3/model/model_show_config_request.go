package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigRequest Request Object
type ShowConfigRequest struct {
}

func (o ShowConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowConfigRequest", string(data)}, " ")
}
