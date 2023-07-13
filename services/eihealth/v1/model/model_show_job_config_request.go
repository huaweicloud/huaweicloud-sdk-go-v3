package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobConfigRequest Request Object
type ShowJobConfigRequest struct {
}

func (o ShowJobConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowJobConfigRequest", string(data)}, " ")
}
