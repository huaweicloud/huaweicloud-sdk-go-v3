package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHostResponse struct {
	// 主机列表

	Result *[]GetHostListInfo `json:"result,omitempty"`
	// 主机信息总数量

	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostResponse struct{}"
	}

	return strings.Join([]string{"ListHostResponse", string(data)}, " ")
}
