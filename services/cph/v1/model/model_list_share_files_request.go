package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShareFilesRequest Request Object
type ListShareFilesRequest struct {

	// 偏移量为一个大于等于0整数，表示查询该偏移量后面的所有的资源数，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页返回的资源个数。取值范围：1~100（默认值为100），一般设置为10、20、50。
	Limit *int32 `json:"limit,omitempty"`

	// 云手机服务器ID列表，多个服务器ID用逗号（,）分隔。
	ServerIds string `json:"server_ids"`

	// 待查询的目录名称， 可以包含大小写字母、数字、“.”、“+”、“-”、“_”、“/”、\"=\"；必须以“/”开头，并且不能只包含“/”；不能包含“../”、“//”等。 示例：/data/data
	Path string `json:"path"`
}

func (o ListShareFilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShareFilesRequest struct{}"
	}

	return strings.Join([]string{"ListShareFilesRequest", string(data)}, " ")
}
