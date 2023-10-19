package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessClientInfosRequest Request Object
type ListAccessClientInfosRequest struct {

	// LakeFormation实例ID。创建实例时自动生成。例如：2180518f-42b8-4947-b20b-adfc53981a25。
	InstanceId string `json:"instance_id"`

	// ID搜索。根据ID进行搜索。
	Id *string `json:"id,omitempty"`

	// 名称关键字搜索。只能包含字母、数字、下划线和中划线，且最大长度为32个字符。
	Name *string `json:"name,omitempty"`

	// 分页查询时的偏移量。默认值为0。最小值为0，最大值为1000。
	Offset int32 `json:"offset"`

	// 分页一页显示数。默认值为10。最小值为1，最大值为1000。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAccessClientInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessClientInfosRequest struct{}"
	}

	return strings.Join([]string{"ListAccessClientInfosRequest", string(data)}, " ")
}
