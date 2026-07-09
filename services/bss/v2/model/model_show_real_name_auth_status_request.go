package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRealNameAuthStatusRequest Request Object
type ShowRealNameAuthStatusRequest struct {
}

func (o ShowRealNameAuthStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealNameAuthStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowRealNameAuthStatusRequest", string(data)}, " ")
}
