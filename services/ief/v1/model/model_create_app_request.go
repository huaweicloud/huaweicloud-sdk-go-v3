package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppRequest Request Object
type CreateAppRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *App `json:"body,omitempty"`
}

func (o CreateAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppRequest struct{}"
	}

	return strings.Join([]string{"CreateAppRequest", string(data)}, " ")
}
