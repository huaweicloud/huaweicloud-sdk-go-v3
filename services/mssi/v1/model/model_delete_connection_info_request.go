package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConnectionInfoRequest Request Object
type DeleteConnectionInfoRequest struct {

	// ID of Connection
	ConnectId string `json:"connect_id"`
}

func (o DeleteConnectionInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnectionInfoRequest struct{}"
	}

	return strings.Join([]string{"DeleteConnectionInfoRequest", string(data)}, " ")
}
