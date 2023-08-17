package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunDeleteDataRequest Request Object
type RunDeleteDataRequest struct {

	// 服务实例的名称，用户创建服务实例时指定。
	ServiceName string `json:"service_name"`

	Body *DeleteParam `json:"body,omitempty"`
}

func (o RunDeleteDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDeleteDataRequest struct{}"
	}

	return strings.Join([]string{"RunDeleteDataRequest", string(data)}, " ")
}
