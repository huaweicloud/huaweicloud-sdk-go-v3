package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CatalogVo 主题目录。
type CatalogVo struct {

	// 中文名称。
	NameCh *string `json:"name_ch,omitempty"`

	// 英文名称。
	NameEn *string `json:"name_en,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 扩展名。
	QualifiedName *string `json:"qualified_name,omitempty"`

	// guid，自动生成。
	Guid *string `json:"guid,omitempty"`

	// 编码。
	Code *string `json:"code,omitempty"`

	// 别名。
	Alias *string `json:"alias,omitempty"`

	Status *BizStatusEnum `json:"status,omitempty"`

	NewBiz *BizVersionManageVo `json:"new_biz,omitempty"`

	// 数据所有者。
	DataOwner *string `json:"data_owner,omitempty"`

	// 数据所有者集合。
	DataOwnerList *string `json:"data_owner_list,omitempty"`

	// 数据域。
	DataDepartment *string `json:"data_department,omitempty"`

	// 路径信息。
	Path *string `json:"path,omitempty"`

	// 层级信息。
	Level *int32 `json:"level,omitempty"`

	// 序号。
	Ordinal *int32 `json:"ordinal,omitempty"`

	// 责任人。
	Owner *string `json:"owner,omitempty"`

	// 父目录ID，木有则为根目录，填写String类型替代Long类型。
	ParentId *string `json:"parent_id,omitempty"`

	// 同层排序，目标节点的ID，填写String类型替代Long类型。
	SwapOrderId *string `json:"swap_order_id,omitempty"`

	// 主题ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 认证ID，自动生成。
	QualifiedId *string `json:"qualified_id,omitempty"`

	// 是否来自公共层。
	FromPublic *bool `json:"from_public,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 拥有子流程的数量，不包括子流程的子流程，前端不传。
	ChildrenNum *int32 `json:"children_num,omitempty"`

	// 下层子目录，只读。
	Children *[]CatalogVo `json:"children,omitempty"`

	// 属性自定义项。
	SelfDefinedFields *[]SelfDefinedFieldVo `json:"self_defined_fields,omitempty"`
}

func (o CatalogVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CatalogVo struct{}"
	}

	return strings.Join([]string{"CatalogVo", string(data)}, " ")
}
