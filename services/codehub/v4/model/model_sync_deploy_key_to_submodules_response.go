package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncDeployKeyToSubmodulesResponse Response Object
type SyncDeployKeyToSubmodulesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SyncDeployKeyToSubmodulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncDeployKeyToSubmodulesResponse struct{}"
	}

	return strings.Join([]string{"SyncDeployKeyToSubmodulesResponse", string(data)}, " ")
}
