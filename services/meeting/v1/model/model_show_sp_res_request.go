package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpResRequest Request Object
type ShowSpResRequest struct {
}

func (o ShowSpResRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpResRequest struct{}"
	}

	return strings.Join([]string{"ShowSpResRequest", string(data)}, " ")
}
