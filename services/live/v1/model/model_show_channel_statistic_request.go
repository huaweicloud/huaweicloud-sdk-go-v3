package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowChannelStatisticRequest Request Object
type ShowChannelStatisticRequest struct {

	// 服务鉴权Token，服务开启鉴权，必须携带Access-Control-Allow-Internal访问服务。
	AccessControlAllowInternal *string `json:"Access-Control-Allow-Internal,omitempty"`

	// 服务鉴权Token，服务开启鉴权，必须携带Access-Control-Allow-External访问服务。
	AccessControlAllowExternal *string `json:"Access-Control-Allow-External,omitempty"`

	// 每页记录数，取值范围[1,100]，默认值10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量。表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	Body *ShowChannelStatisticReq `json:"body,omitempty"`
}

func (o ShowChannelStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowChannelStatisticRequest struct{}"
	}

	return strings.Join([]string{"ShowChannelStatisticRequest", string(data)}, " ")
}
