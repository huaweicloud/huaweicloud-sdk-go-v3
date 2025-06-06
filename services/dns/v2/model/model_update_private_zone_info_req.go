package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePrivateZoneInfoReq struct {

	// 域名的描述信息。长度不超过255个字符。
	Description *string `json:"description,omitempty"`

	// 管理该域名的管理员邮箱，用于生成该域名的SOA记录。
	Email *string `json:"email,omitempty"`

	// 用于填写默认生成的SOA记录中有效缓存时间，以秒为单位。
	Ttl *int32 `json:"ttl,omitempty"`
}

func (o UpdatePrivateZoneInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateZoneInfoReq struct{}"
	}

	return strings.Join([]string{"UpdatePrivateZoneInfoReq", string(data)}, " ")
}
