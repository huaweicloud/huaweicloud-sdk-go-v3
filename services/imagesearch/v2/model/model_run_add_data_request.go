package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunAddDataRequest struct {

	// 服务实例的名称，用户创建服务实例时指定。
	ServiceName string `json:"service_name"`

	Body *AddDataParam `json:"body,omitempty"`
}

func (o RunAddDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAddDataRequest struct{}"
	}

	return strings.Join([]string{"RunAddDataRequest", string(data)}, " ")
}
