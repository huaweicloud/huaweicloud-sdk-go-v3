package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebTamperProtectionContainerResponse Response Object
type ListWebTamperProtectionContainerResponse struct {

	// **参数解释**: 防护配置关联的容器信息总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 防护配置关联的容器信息列表 **取值范围**: 取值0-200个WebTamperProtectionContainerResponseInfo对象
	DataList       *[]WebTamperProtectionContainerResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o ListWebTamperProtectionContainerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebTamperProtectionContainerResponse struct{}"
	}

	return strings.Join([]string{"ListWebTamperProtectionContainerResponse", string(data)}, " ")
}
