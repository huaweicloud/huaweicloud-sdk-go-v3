package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationEndpointAttributesResponseBodyAttributes 属性列表。
type ListApplicationEndpointAttributesResponseBodyAttributes struct {

	// 设备是否可用。
	Enabled string `json:"enabled"`

	// 设备token。
	Token string `json:"token"`

	// 用户数据。
	UserData string `json:"user_data"`
}

func (o ListApplicationEndpointAttributesResponseBodyAttributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationEndpointAttributesResponseBodyAttributes struct{}"
	}

	return strings.Join([]string{"ListApplicationEndpointAttributesResponseBodyAttributes", string(data)}, " ")
}
