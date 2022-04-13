package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowSpResRequest struct {
}

func (o ShowSpResRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpResRequest struct{}"
	}

	return strings.Join([]string{"ShowSpResRequest", string(data)}, " ")
}
