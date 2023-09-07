package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSingleConnectionRequest Request Object
type ShowSingleConnectionRequest struct {

	// ID of Connection
	ConnectId string `json:"connect_id"`
}

func (o ShowSingleConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowSingleConnectionRequest", string(data)}, " ")
}
