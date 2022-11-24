package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDefaultConfigRequest struct {
}

func (o ShowDefaultConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowDefaultConfigRequest", string(data)}, " ")
}
