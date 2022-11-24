package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBaremetalServerTagsRequest struct {
	ServerId string `json:"server_id"`
}

func (o ShowBaremetalServerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaremetalServerTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowBaremetalServerTagsRequest", string(data)}, " ")
}
