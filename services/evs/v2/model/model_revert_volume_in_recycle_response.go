package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevertVolumeInRecycleResponse Response Object
type RevertVolumeInRecycleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RevertVolumeInRecycleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevertVolumeInRecycleResponse struct{}"
	}

	return strings.Join([]string{"RevertVolumeInRecycleResponse", string(data)}, " ")
}
