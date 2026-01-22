package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateInfoRequest Request Object
type ShowCertificateInfoRequest struct {

	// 播放域名
	PlayDomain *string `json:"play_domain,omitempty"`

	// 偏移量，表示从此偏移量开始查询
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，取值范围[1,100]
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowCertificateInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificateInfoRequest", string(data)}, " ")
}
