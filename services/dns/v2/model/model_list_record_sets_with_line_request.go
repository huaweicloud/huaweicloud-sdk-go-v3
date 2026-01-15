package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecordSetsWithLineRequest Request Object
type ListRecordSetsWithLineRequest struct {

	// **参数解释：** 记录集所属域名的类型。 **约束限制：** 不涉及。 **取值范围：** - public：表示查询公网域名的记录集。 - private：表示查询内网域名的记录集。  **默认取值：** public。
	ZoneType *string `json:"zone_type,omitempty"`

	// **参数解释：** 分页查询的起始资源ID。 - 查询第一页时，设置为空。 - 查询下一页时，设置为上一页最后一条资源的ID。  **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释：** 分页查询时配置每页返回的资源个数。 **约束限制：** 不涉及。 **取值范围：** 0~500。 **默认取值：** 500
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。 **约束限制：** 当设置marker不为空时，以marker为分页起始标识，offset不生效。 **取值范围：** 0~2147483647。 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 域名的ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneId *string `json:"zone_id,omitempty"`

	// **参数解释：** 解析线路ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	LineId *string `json:"line_id,omitempty"`

	// **参数解释：** 记录集的标签，包括标签键和标签值。 取值格式：key1,value1|key2,value2。 **约束限制：** - 多个标签之间用“|”分开，每个标签的键值用英文逗号“,”相隔。 - 多个标签之间为“与”的关系。 - 搜索模式为精确搜索。如果资源标签值value是以&ast;开头时，则按照&ast;后面的值全模糊匹配。  **取值范围：** 最多可以查询20个标签。 **默认取值：** 不涉及。
	Tags *string `json:"tags,omitempty"`

	// **参数解释：** 记录集状态。 **约束限制：** 不涉及。 **取值范围：** - ACTIVE：正常 - PENDING_CREATE：创建中 - PENDING_UPDATE：更新中 - PENDING_DELETE：删除中 - PENDING_FREEZE：冻结中 - FREEZE：冻结 - ILLEGAL：违规冻结 - POLICE：公安冻结 - PENDING_DISABLE：暂停中 - DISABLE：暂停 - ERROR：失败  **默认取值：** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 记录集的类型。 **约束限制：** 不涉及。 **取值范围：** - A：将域名解析到指定的IPv4地址。 - AAAA：将域名解析到指定的IPv6地址。 - MX：指定域名对应的邮件服务器。 - CNAME：将域名解析到另一域名，或者多个域名映射到同一域名上。 - TXT：用于对域名进行标识和说明。 - SRV：用于记录某台服务器对外提供了哪些服务。 - NS：指定域名的权威DNS服务器。 - SOA：提供域名的基本信息和权威服务器的详细信息。 - CAA：指定为域名颁发HTTPS证书的授权CA机构，仅支持公网域名解析。 - PTR：指定IP地址反向解析记录，仅支持内网域名解析。  **默认取值：** 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 待查询的记录集的域名中包含此name。 搜索模式默认为模糊搜索。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 待查询的记录集ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 待查询的记录集的值。 搜索模式默认为模糊搜索。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Records *string `json:"records,omitempty"`

	// **参数解释：** 查询结果中记录集列表的排序字段。 **约束限制：** 不涉及。 **取值范围：** - name：记录集名称 - type：记录集类型 - created_at：创建时间 - updated_at：更新时间  **默认取值：** created_at
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释：** 查询结果中记录集列表的排序方式。 **约束限制：** 不涉及。 **取值范围：** - desc：降序排序 - asc：升序排序  **默认取值：** desc
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释：** 健康检查ID。 搜索模式默认为模糊搜索。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	HealthCheckId *string `json:"health_check_id,omitempty"`

	// **参数解释：** 查询条件搜索模式。 **约束限制：** 不涉及。 **取值范围：** - like：模糊搜索 - equal：精确搜索  **默认取值：** name默认模糊搜索，其他默认精确搜索。
	SearchMode *string `json:"search_mode,omitempty"`
}

func (o ListRecordSetsWithLineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordSetsWithLineRequest struct{}"
	}

	return strings.Join([]string{"ListRecordSetsWithLineRequest", string(data)}, " ")
}
