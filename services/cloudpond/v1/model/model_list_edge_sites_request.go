package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeSitesRequest Request Object
type ListEdgeSitesRequest struct {

	// 每页的数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识
	Marker *string `json:"marker,omitempty"`

	// 排序字段
	SortKey *[]string `json:"sort_key,omitempty"`

	// 排序方向，取值范围： - desc：降序 - acs：升序
	SortDir *[]string `json:"sort_dir,omitempty"`

	// 企业项目ID。可以使用该字段过滤某个企业项目下的边缘小站。 最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。 若需要查询当前用户所有企业项目绑定的边缘小站，请传参all_granted_eps。 不传则查询全部。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据边缘小站ID查询，支持排序
	Id *[]string `json:"id,omitempty"`

	// 根据边缘小站名称查询（精确），支持排序
	Name *[]string `json:"name,omitempty"`

	// 根据边缘可用区ID查询
	AvailabilityZoneId *[]string `json:"availability_zone_id,omitempty"`

	// 根据边缘小站部署状态查询
	Status *[]SiteStatus `json:"status,omitempty"`
}

func (o ListEdgeSitesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeSitesRequest struct{}"
	}

	return strings.Join([]string{"ListEdgeSitesRequest", string(data)}, " ")
}
