package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecreateServerImageRequest Request Object
type RecreateServerImageRequest struct {

	// 镜像实例唯一标识。
	ServerId string `json:"server_id"`

	Body *CreateServerImageReq `json:"body,omitempty"`
}

func (o RecreateServerImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecreateServerImageRequest struct{}"
	}

	return strings.Join([]string{"RecreateServerImageRequest", string(data)}, " ")
}
