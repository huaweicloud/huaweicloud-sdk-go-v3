package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckPasswordRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	Body *CheckPasswordRequestBody `json:"body,omitempty"`
}

func (o CheckPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPasswordRequest struct{}"
	}

	return strings.Join([]string{"CheckPasswordRequest", string(data)}, " ")
}
