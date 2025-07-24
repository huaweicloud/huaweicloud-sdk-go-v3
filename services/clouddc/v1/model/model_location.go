package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Location 物理服务器部署位置信息
type Location struct {

	// 服务器所在的机房名称
	Dc *string `json:"dc,omitempty"`

	// 服务器所在的iRack名称
	Rack *string `json:"rack,omitempty"`

	// 服务器所在的U位
	Unit *string `json:"unit,omitempty"`
}

func (o Location) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Location struct{}"
	}

	return strings.Join([]string{"Location", string(data)}, " ")
}
