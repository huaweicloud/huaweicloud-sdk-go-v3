package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckLaunchResponse Response Object
type CheckLaunchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckLaunchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckLaunchResponse struct{}"
	}

	return strings.Join([]string{"CheckLaunchResponse", string(data)}, " ")
}
