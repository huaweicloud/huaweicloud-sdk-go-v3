package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSystemConfigsResponse Response Object
type ListSystemConfigsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 系统配置信息列表
	DataList       *[]string `json:"data_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSystemConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListSystemConfigsResponse", string(data)}, " ")
}
