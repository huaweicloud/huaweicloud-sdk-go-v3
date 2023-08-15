package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceRemarkRequest Request Object
type UpdateInstanceRemarkRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId string `json:"instance_id"`

	Body *UpdateInstanceRemarkRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceRemarkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceRemarkRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRemarkRequest", string(data)}, " ")
}
