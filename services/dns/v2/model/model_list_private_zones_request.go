package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateZonesRequest Request Object
type ListPrivateZonesRequest struct {

	// 待查询的域名的类型。  取值范围：private。
	Type string `json:"type"`

	// 每页返回的资源个数。  取值范围：0~500  取值一般为10，20，50。默认值为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的资源ID，为空时为查询第一页。  默认值为空。
	Marker *string `json:"marker,omitempty"`

	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。  取值范围：0~2147483647  默认值为0。  当前设置marker不为空时，以marker为分页起始标识。
	Offset *int32 `json:"offset,omitempty"`

	// 资源标签。
	Tags *string `json:"tags,omitempty"`

	// 域名。  搜索模式默认为模糊搜索。
	Name *string `json:"name,omitempty"`

	// 域名ID。
	Id *string `json:"id,omitempty"`

	// 资源状态。
	Status *string `json:"status,omitempty"`

	// 查询条件搜索模式。  取值范围：  like：模糊搜索 equal：精确搜索
	SearchMode *string `json:"search_mode,omitempty"`

	// 查询结果中域名列表的排序字段。  取值范围为：  name：域名 created_at：创建时间 updated_at：更新时间 默认值为空，表示不排序。
	SortKey *string `json:"sort_key,omitempty"`

	// 查询结果中域名列表的排序方式。  取值范围：  desc：降序排序 asc：升序排序 默认值为空，表示不排序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 域名关联的企业项目ID，长度不超过36个字符。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 关联VPC的ID。
	RouterId *string `json:"router_id,omitempty"`
}

func (o ListPrivateZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateZonesRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateZonesRequest", string(data)}, " ")
}
