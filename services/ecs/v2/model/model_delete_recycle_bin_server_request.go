package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecycleBinServerRequest Request Object
type DeleteRecycleBinServerRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o DeleteRecycleBinServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecycleBinServerRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecycleBinServerRequest", string(data)}, " ")
}
