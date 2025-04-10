package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRansomwareProtectionNodesResponse Response Object
type ListRansomwareProtectionNodesResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 查询勒索防护服务器列表
	DataList       *[]ProtectionServerInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListRansomwareProtectionNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRansomwareProtectionNodesResponse struct{}"
	}

	return strings.Join([]string{"ListRansomwareProtectionNodesResponse", string(data)}, " ")
}
