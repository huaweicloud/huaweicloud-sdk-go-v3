package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTestsuiteInfosUsingRequest Request Object
type DeleteTestsuiteInfosUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	Body *DeleteTaskParams `json:"body,omitempty"`
}

func (o DeleteTestsuiteInfosUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTestsuiteInfosUsingRequest struct{}"
	}

	return strings.Join([]string{"DeleteTestsuiteInfosUsingRequest", string(data)}, " ")
}
