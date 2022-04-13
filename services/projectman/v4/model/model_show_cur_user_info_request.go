package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCurUserInfoRequest struct {
}

func (o ShowCurUserInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCurUserInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowCurUserInfoRequest", string(data)}, " ")
}
