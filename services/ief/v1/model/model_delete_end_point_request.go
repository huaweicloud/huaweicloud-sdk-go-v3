package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEndPointRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	// 端点ID
	EndpointId string `json:"endpoint_id" xml:"endpoint_id"`
}

func (o DeleteEndPointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEndPointRequest struct{}"
	}

	return strings.Join([]string{"DeleteEndPointRequest", string(data)}, " ")
}
