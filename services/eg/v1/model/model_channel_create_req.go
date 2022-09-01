package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChannelCreateReq struct {

	// 通道名称，租户下唯一，由字母，数字，点，下划线和中划线组成，必须字母或数字开头，不能是default
	Name *string `json:"name,omitempty" xml:"name"`

	// 通道描述
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o ChannelCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelCreateReq struct{}"
	}

	return strings.Join([]string{"ChannelCreateReq", string(data)}, " ")
}
