package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrepayResizeRequest Request Object
type CreatePrepayResizeRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *InstanceChangeOrderReq `json:"body,omitempty"`
}

func (o CreatePrepayResizeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrepayResizeRequest struct{}"
	}

	return strings.Join([]string{"CreatePrepayResizeRequest", string(data)}, " ")
}
