package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWebTamperProtectionDirectoryResponse Response Object
type ListWebTamperProtectionDirectoryResponse struct {

	// **参数解释**: 防护配置关联的容器的防护目录总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 防护配置关联的容器的防护目录列表 **取值范围**: 取值0-200个WebTamperProtectionDirectoryResponseInfo对象
	DataList       *[]WebTamperProtectionDirectoryResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o ListWebTamperProtectionDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWebTamperProtectionDirectoryResponse struct{}"
	}

	return strings.Join([]string{"ListWebTamperProtectionDirectoryResponse", string(data)}, " ")
}
