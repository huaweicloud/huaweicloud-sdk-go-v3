package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevertRecycleBinServerRequest Request Object
type RevertRecycleBinServerRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o RevertRecycleBinServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevertRecycleBinServerRequest struct{}"
	}

	return strings.Join([]string{"RevertRecycleBinServerRequest", string(data)}, " ")
}
