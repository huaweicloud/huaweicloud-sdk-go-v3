package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBareMetalServerDetailsRequest struct {
	// 裸金属服务器ID

	ServerId string `json:"server_id"`
}

func (o ListBareMetalServerDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBareMetalServerDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListBareMetalServerDetailsRequest", string(data)}, " ")
}
