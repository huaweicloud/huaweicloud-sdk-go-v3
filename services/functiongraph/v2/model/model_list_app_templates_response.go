package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppTemplatesResponse Response Object
type ListAppTemplatesResponse struct {

	// 函数应用程序模板列表
	Templates *[]ListAppTemplatesResult `json:"templates,omitempty"`

	// 下次读取位置
	NextMarker *int64 `json:"next_marker,omitempty"`

	// 应用程序模板总数
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAppTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAppTemplatesResponse", string(data)}, " ")
}
