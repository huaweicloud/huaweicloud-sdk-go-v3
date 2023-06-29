package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSourceIpResponse Response Object
type ShowSourceIpResponse struct {

	// 源站信息列表
	SourceIp *[]IpsItem `json:"source_ip,omitempty"`

	// 回源Ip最后更新时间
	LastModify     *int64 `json:"last_modify,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSourceIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSourceIpResponse struct{}"
	}

	return strings.Join([]string{"ShowSourceIpResponse", string(data)}, " ")
}
