package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppDetailRequest Request Object
type ShowAppDetailRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 应用模板ID
	AppId string `json:"app_id"`
}

func (o ShowAppDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowAppDetailRequest", string(data)}, " ")
}
