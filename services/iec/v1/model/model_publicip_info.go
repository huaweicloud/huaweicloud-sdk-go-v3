package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性公网IP信息
type PublicipInfo struct {
	//   IP版本的信息

	IpVersion *int32 `json:"ip_version,omitempty"`
	// 弹性公网IP

	PublicipAddress *string `json:"publicip_address,omitempty"`
	// 弹性公网IP的ID。

	PublicipId *string `json:"publicip_id,omitempty"`
	// 功能说明：弹性公网IP的类型

	PublicipType *string `json:"publicip_type,omitempty"`
}

func (o PublicipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicipInfo struct{}"
	}

	return strings.Join([]string{"PublicipInfo", string(data)}, " ")
}
