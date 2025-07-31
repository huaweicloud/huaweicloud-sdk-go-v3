package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFileHostsResponse Response Object
type ListFileHostsResponse struct {

	// 变更云服务器数量
	TotalNum *int32 `json:"total_num,omitempty"`

	// 变更云服务器信息列表
	DataList       *[]FileHostResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListFileHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFileHostsResponse struct{}"
	}

	return strings.Join([]string{"ListFileHostsResponse", string(data)}, " ")
}
