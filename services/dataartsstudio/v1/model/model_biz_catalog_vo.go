package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BizCatalogVo 流程架构目录。
type BizCatalogVo struct {

	// 流程名称。
	Name string `json:"name"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 对应资产中ID。
	Guid *string `json:"guid,omitempty"`

	// 责任人。
	Owner string `json:"owner"`

	// 父目录ID，没有则为根目录。填写String类型替代Long类型。
	ParentId *string `json:"parent_id,omitempty"`

	// 上个节点ID，没有则为首节点。填写String类型替代Long类型。
	PrevId *string `json:"prev_id,omitempty"`

	// 下个节点ID，没有则为尾节点。填写String类型替代Long类型。
	NextId *string `json:"next_id,omitempty"`

	// 创建时传空，更新时必填。填写String类型替代Long类型。
	Id string `json:"id"`

	// 认证ID，自动生成。
	QualifiedId *string `json:"qualified_id,omitempty"`

	// 创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 拥有业务指标数量，前端不传。
	BizmetricNum *int32 `json:"bizmetric_num,omitempty"`

	// 拥有子流程的数量，不包括子流程的子流程。
	ChildrenNum *int32 `json:"children_num,omitempty"`

	// 下层子目录，只读。
	Children *[]BizCatalogVo `json:"children,omitempty"`
}

func (o BizCatalogVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BizCatalogVo struct{}"
	}

	return strings.Join([]string{"BizCatalogVo", string(data)}, " ")
}
