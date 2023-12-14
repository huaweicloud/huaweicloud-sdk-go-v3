package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchClusterTagsResponse Response Object
type SwitchClusterTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchClusterTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchClusterTagsResponse struct{}"
	}

	return strings.Join([]string{"SwitchClusterTagsResponse", string(data)}, " ")
}
