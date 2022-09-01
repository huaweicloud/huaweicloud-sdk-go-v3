package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ReinstallServerWithCloudInitRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id" xml:"server_id"`

	Body *ReinstallServerWithCloudInitRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ReinstallServerWithCloudInitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerWithCloudInitRequest struct{}"
	}

	return strings.Join([]string{"ReinstallServerWithCloudInitRequest", string(data)}, " ")
}
