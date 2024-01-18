package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckLaunchRequest Request Object
type CheckLaunchRequest struct {
}

func (o CheckLaunchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckLaunchRequest struct{}"
	}

	return strings.Join([]string{"CheckLaunchRequest", string(data)}, " ")
}
