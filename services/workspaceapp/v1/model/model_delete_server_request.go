package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerRequest Request Object
type DeleteServerRequest struct {

	// 服务器唯一标识。
	ServerId string `json:"server_id"`
}

func (o DeleteServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerRequest", string(data)}, " ")
}
