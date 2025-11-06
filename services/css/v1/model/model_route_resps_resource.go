package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RouteRespsResource struct {

	// ip地址。
	IpAddress *string `json:"ipAddress,omitempty"`

	// 子网掩码。
	IpNetMask *string `json:"ipNetMask,omitempty"`

	// 更新时间。
	UpdateAt *string `json:"updateAt,omitempty"`
}

func (o RouteRespsResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteRespsResource struct{}"
	}

	return strings.Join([]string{"RouteRespsResource", string(data)}, " ")
}
