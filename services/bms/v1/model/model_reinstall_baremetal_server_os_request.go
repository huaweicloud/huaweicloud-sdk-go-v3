package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ReinstallBaremetalServerOsRequest struct {

	// 裸金属服务器ID
	ServerId string `json:"server_id"`

	Body *OsReinstallBody `json:"body,omitempty"`
}

func (o ReinstallBaremetalServerOsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallBaremetalServerOsRequest struct{}"
	}

	return strings.Join([]string{"ReinstallBaremetalServerOsRequest", string(data)}, " ")
}
