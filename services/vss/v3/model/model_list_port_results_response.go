package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortResultsResponse Response Object
type ListPortResultsResponse struct {

	// 端口总数
	Total *int32 `json:"total,omitempty"`

	// 端口信息列表
	Data           *[]PortItem `json:"data,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListPortResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortResultsResponse struct{}"
	}

	return strings.Join([]string{"ListPortResultsResponse", string(data)}, " ")
}
