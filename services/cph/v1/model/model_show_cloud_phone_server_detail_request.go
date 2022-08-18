package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCloudPhoneServerDetailRequest struct {

	// 服务器id。
	ServerId string `json:"server_id"`
}

func (o ShowCloudPhoneServerDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudPhoneServerDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowCloudPhoneServerDetailRequest", string(data)}, " ")
}
