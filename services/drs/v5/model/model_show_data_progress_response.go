package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataProgressResponse Response Object
type ShowDataProgressResponse struct {

	// 数据加工规则响应体
	DataProcessInfo *[]DataProcessInfo `json:"data_process_info,omitempty"`

	// 数据加工规则总数目
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDataProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowDataProgressResponse", string(data)}, " ")
}
