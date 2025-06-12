package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecycleBinServerRequest Request Object
type ShowRecycleBinServerRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ShowRecycleBinServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecycleBinServerRequest struct{}"
	}

	return strings.Join([]string{"ShowRecycleBinServerRequest", string(data)}, " ")
}
