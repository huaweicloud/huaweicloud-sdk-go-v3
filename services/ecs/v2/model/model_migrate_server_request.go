package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateServerRequest Request Object
type MigrateServerRequest struct {

	// 云服务器ID。
	ServerId string `json:"server_id"`

	Body *MigrateServerRequestBody `json:"body,omitempty"`
}

func (o MigrateServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateServerRequest struct{}"
	}

	return strings.Join([]string{"MigrateServerRequest", string(data)}, " ")
}
