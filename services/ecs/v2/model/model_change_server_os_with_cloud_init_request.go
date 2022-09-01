package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeServerOsWithCloudInitRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id" xml:"server_id"`

	Body *ChangeServerOsWithCloudInitRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ChangeServerOsWithCloudInitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithCloudInitRequest struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithCloudInitRequest", string(data)}, " ")
}
