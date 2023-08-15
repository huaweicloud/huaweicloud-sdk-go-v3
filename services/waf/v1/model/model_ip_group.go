package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpGroup struct {

	// Ip地址组id，在新增Ip地址组时系统自动生成的唯一标识
	Id *string `json:"id,omitempty"`

	// Ip地址组名
	Name *string `json:"name,omitempty"`

	// Ip地址组中包含Ip/Ip段的数量
	Size *int64 `json:"size,omitempty"`
}

func (o IpGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroup struct{}"
	}

	return strings.Join([]string{"IpGroup", string(data)}, " ")
}
