package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostPayResizeOrderRequest Request Object
type CreatePostPayResizeOrderRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *ResizeInstanceReq `json:"body,omitempty"`
}

func (o CreatePostPayResizeOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPayResizeOrderRequest struct{}"
	}

	return strings.Join([]string{"CreatePostPayResizeOrderRequest", string(data)}, " ")
}
