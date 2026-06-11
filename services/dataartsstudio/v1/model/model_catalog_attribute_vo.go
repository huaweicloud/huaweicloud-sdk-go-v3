package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CatalogAttributeVo struct {

	// 路径。用“/”作为分隔，如：运营领域/云学院/云学院培训方案。
	Path string `json:"path"`

	// 资产名称。
	QualifiedName string `json:"qualifiedName"`

	// 主题所属层级。
	Level string `json:"level"`

	// 名称。
	Name string `json:"name"`

	// 英文名称。
	NameEng string `json:"nameEng"`

	// 别名。
	Alias *string `json:"alias,omitempty"`

	// 描述。
	Description string `json:"description"`

	// 数据主体。
	DataOwner string `json:"dataOwner"`

	// 责任人。
	Owner *string `json:"owner,omitempty"`

	// 数据主体列表。
	DataOwnerList []string `json:"dataOwnerList"`

	// 创建时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *string `json:"createTime,omitempty"`

	// 创建人。
	CreateBy *string `json:"createBy,omitempty"`

	// 更新时间，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *string `json:"updateTime,omitempty"`

	// 更新人。
	UpdateBy *string `json:"updateBy,omitempty"`

	Parent *CatalogAttributeVoParent `json:"parent,omitempty"`

	// 父节点ID。
	ParentId *string `json:"parentId,omitempty"`

	// 是否为L1层。主题域分组。
	L1 *bool `json:"l1,omitempty"`

	// 是否为L2层。主题域。
	L2 *bool `json:"l2,omitempty"`

	// 是否为L3层。业务对象。
	L3 *bool `json:"l3,omitempty"`

	// 顺序编号。主题返回时根据此编号由小到大排序。
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 租户ID。获取方式参考此接口的路径参数“project_id”获取。
	TenantId *string `json:"tenantId,omitempty"`

	// 自定义项。主题的自定义属性项。
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`
}

func (o CatalogAttributeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogAttributeVo struct{}"
	}

	return strings.Join([]string{"CatalogAttributeVo", string(data)}, " ")
}
