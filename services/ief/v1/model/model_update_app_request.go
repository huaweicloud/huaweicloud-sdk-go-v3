package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppRequest Request Object
type UpdateAppRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 应用模板ID
	AppId string `json:"app_id"`

	Body *UpdateAppBody `json:"body,omitempty"`
}

func (o UpdateAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppRequest", string(data)}, " ")
}
