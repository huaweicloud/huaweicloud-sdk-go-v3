package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteIpListOption 批量删除IP地址组中的IP。
type BatchDeleteIpListOption struct {

	// IP列表。
	IpList *[]IpGroupIp `json:"ip_list,omitempty"`
}

func (o BatchDeleteIpListOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpListOption struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpListOption", string(data)}, " ")
}
