package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunCheckDataRequest struct {

	// 服务实例的名称，用户创建服务实例时指定。
	ServiceName string `json:"service_name"`

	Body *CheckParam `json:"body,omitempty"`
}

func (o RunCheckDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCheckDataRequest struct{}"
	}

	return strings.Join([]string{"RunCheckDataRequest", string(data)}, " ")
}
