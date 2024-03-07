package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Port 私有IP信息。
type Port struct {

	// 私有IP唯一标识
	Id *string `json:"id,omitempty"`

	// 私有IP地址
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o Port) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Port struct{}"
	}

	return strings.Join([]string{"Port", string(data)}, " ")
}
