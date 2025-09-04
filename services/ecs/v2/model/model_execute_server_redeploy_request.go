package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteServerRedeployRequest Request Object
type ExecuteServerRedeployRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`
}

func (o ExecuteServerRedeployRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteServerRedeployRequest struct{}"
	}

	return strings.Join([]string{"ExecuteServerRedeployRequest", string(data)}, " ")
}
