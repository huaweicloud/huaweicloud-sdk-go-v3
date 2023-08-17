package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunUpdateDataRequest Request Object
type RunUpdateDataRequest struct {

	// 服务实例的名称，用户创建服务实例时指定。
	ServiceName string `json:"service_name"`

	Body *UpdateParam `json:"body,omitempty"`
}

func (o RunUpdateDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunUpdateDataRequest struct{}"
	}

	return strings.Join([]string{"RunUpdateDataRequest", string(data)}, " ")
}
