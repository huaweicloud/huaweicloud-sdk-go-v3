package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImageServerRequest Request Object
type UpdateImageServerRequest struct {

	// 镜像实例唯一标识。
	ServerId string `json:"server_id"`

	Body *UpdateImageServerReq `json:"body,omitempty"`
}

func (o UpdateImageServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageServerRequest struct{}"
	}

	return strings.Join([]string{"UpdateImageServerRequest", string(data)}, " ")
}
