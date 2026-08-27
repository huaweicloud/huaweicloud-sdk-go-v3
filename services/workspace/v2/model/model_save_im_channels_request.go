package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveImChannelsRequest Request Object
type SaveImChannelsRequest struct {
	Body *SaveImChannelsReq `json:"body,omitempty"`
}

func (o SaveImChannelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveImChannelsRequest struct{}"
	}

	return strings.Join([]string{"SaveImChannelsRequest", string(data)}, " ")
}
