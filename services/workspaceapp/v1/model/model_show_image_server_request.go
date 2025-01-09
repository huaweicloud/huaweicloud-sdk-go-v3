package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageServerRequest Request Object
type ShowImageServerRequest struct {

	// 镜像实例唯一标识。
	ServerId string `json:"server_id"`
}

func (o ShowImageServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageServerRequest struct{}"
	}

	return strings.Join([]string{"ShowImageServerRequest", string(data)}, " ")
}
