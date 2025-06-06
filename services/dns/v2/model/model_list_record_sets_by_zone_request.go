package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecordSetsByZoneRequest Request Object
type ListRecordSetsByZoneRequest struct {

	// 域名ID。
	ZoneId string `json:"zone_id"`

	// 查询条件搜索模式。  取值范围：  like：模糊搜索 equal：精确搜索
	SearchMode *string `json:"search_mode,omitempty"`

	// 分页查询起始的资源ID，为空时为查询第一页。  默认值为空。
	Marker *string `json:"marker,omitempty"`

	// 每页返回的资源个数。  取值范围：0~500  取值一般为10，20，50。默认值为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。  取值范围：0~2147483647  默认值为0。  当前设置marker不为空时，以marker为分页起始标识。
	Offset *int32 `json:"offset,omitempty"`

	// 资源标签。  取值格式：key1,value1|key2,value2  多个标签之间用“|”分开，每个标签的键值用英文逗号“,”相隔。
	Tags *string `json:"tags,omitempty"`

	// 待查询的记录集的状态。 取值范围：ACTIVE、ERROR、DISABLE、FREEZE、PENDING_CREATE、PENDING_UPDATE、PENDING_DELETE
	Status *string `json:"status,omitempty"`

	// 待查询的记录集的类型。 公网域名场景的记录类型: A、AAAA、MX、CNAME、TXT、NS、SRV、CAA。 内网域名场景的记录类型: A、AAAA、MX、CNAME、TXT、SRV。
	Type *string `json:"type,omitempty"`

	// 待查询的记录集的域名中包含此name。  搜索模式默认为模糊搜索。  默认值为空。
	Name *string `json:"name,omitempty"`

	// 待查询的记录集ID。
	Id *string `json:"id,omitempty"`

	// 查询结果中记录集列表的排序字段。  取值范围为：  name：记录集名称 type：记录集类型 默认值为空，表示不排序。
	SortKey *string `json:"sort_key,omitempty"`

	// 查询结果中记录集列表的排序方式。  取值范围：  desc：降序排序 asc：升序排序 默认值为空，表示不排序。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListRecordSetsByZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsByZoneRequest struct{}"
	}

	return strings.Join([]string{"ListRecordSetsByZoneRequest", string(data)}, " ")
}
