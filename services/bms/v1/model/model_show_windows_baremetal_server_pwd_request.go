package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowWindowsBaremetalServerPwdRequest struct {

	// 裸金属服务器ID
	ServerId string `json:"server_id"`
}

func (o ShowWindowsBaremetalServerPwdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWindowsBaremetalServerPwdRequest struct{}"
	}

	return strings.Join([]string{"ShowWindowsBaremetalServerPwdRequest", string(data)}, " ")
}
