package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachImageServerAppRequest Request Object
type AttachImageServerAppRequest struct {

	// 镜像实例唯一标识。
	ServerId string `json:"server_id"`

	Body *AttachServerAppReq `json:"body,omitempty"`
}

func (o AttachImageServerAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachImageServerAppRequest struct{}"
	}

	return strings.Join([]string{"AttachImageServerAppRequest", string(data)}, " ")
}
