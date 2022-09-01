package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListWhitelistsRequest struct {

	// 分页查询中每页的白名单个数
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 分页查询的起始的资源id，表示上一页最后一条查询记录的白名单的id。不指定时表示查询第一页。
	Marker *string `json:"marker,omitempty" xml:"marker"`

	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。
	PageReverse *bool `json:"page_reverse,omitempty" xml:"page_reverse"`

	// 白名单ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 是否开启访问控制开关。true：打开false：关闭
	EnableWhitelist *bool `json:"enable_whitelist,omitempty" xml:"enable_whitelist"`

	// 白名单关联的监听器ID。
	ListenerId *string `json:"listener_id,omitempty" xml:"listener_id"`

	// 白名单IP的字符串。
	Whitelist *string `json:"whitelist,omitempty" xml:"whitelist"`
}

func (o ListWhitelistsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWhitelistsRequest struct{}"
	}

	return strings.Join([]string{"ListWhitelistsRequest", string(data)}, " ")
}
