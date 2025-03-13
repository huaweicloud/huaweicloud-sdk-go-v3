package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumeInfoResponse Response Object
type ListVolumeInfoResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListVolumeInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeInfoResponse struct{}"
	}

	return strings.Join([]string{"ListVolumeInfoResponse", string(data)}, " ")
}
