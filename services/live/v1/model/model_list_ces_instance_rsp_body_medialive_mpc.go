package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCesInstanceRspBodyMedialiveMpc 媒体直播转码服务
type ListCesInstanceRspBodyMedialiveMpc struct {

	// 频道id
	Id *string `json:"id,omitempty"`

	// 频道名
	Name *string `json:"name,omitempty"`
}

func (o ListCesInstanceRspBodyMedialiveMpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCesInstanceRspBodyMedialiveMpc struct{}"
	}

	return strings.Join([]string{"ListCesInstanceRspBodyMedialiveMpc", string(data)}, " ")
}
