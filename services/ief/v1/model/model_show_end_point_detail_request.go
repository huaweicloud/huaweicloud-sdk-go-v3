package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEndPointDetailRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 端点ID
	EndpointId string `json:"endpoint_id"`
}

func (o ShowEndPointDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEndPointDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowEndPointDetailRequest", string(data)}, " ")
}
