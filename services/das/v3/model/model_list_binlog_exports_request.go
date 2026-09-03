package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBinlogExportsRequest Request Object
type ListBinlogExportsRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	// 当前页码
	CurPage *int32 `json:"cur_page,omitempty"`

	// 每页数量
	PerPage *int32 `json:"per_page,omitempty"`
}

func (o ListBinlogExportsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBinlogExportsRequest struct{}"
	}

	return strings.Join([]string{"ListBinlogExportsRequest", string(data)}, " ")
}
