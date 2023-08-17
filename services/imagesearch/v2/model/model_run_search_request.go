package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSearchRequest Request Object
type RunSearchRequest struct {

	// 服务实例的名称，用户创建服务实例时指定。
	ServiceName string `json:"service_name"`

	Body *SearchParam `json:"body,omitempty"`
}

func (o RunSearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSearchRequest struct{}"
	}

	return strings.Join([]string{"RunSearchRequest", string(data)}, " ")
}
