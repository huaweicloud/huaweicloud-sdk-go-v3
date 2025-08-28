package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteIpListOption **参数解释**：批量删除IP地址项的请求参数。  **约束限制**：不涉及
type BatchDeleteIpListOption struct {

	// **参数解释**：IP地址组中IP列表的IP地址信息。  **约束限制**：不涉及
	IpList *[]IpGroupIp `json:"ip_list,omitempty"`
}

func (o BatchDeleteIpListOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpListOption struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpListOption", string(data)}, " ")
}
