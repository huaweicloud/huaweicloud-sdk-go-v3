package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogTreeResponse Response Object
type ListCatalogTreeResponse struct {

	// 流程名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 对应资产中id
	Guid *string `json:"guid,omitempty"`

	// 责任人
	Owner *string `json:"owner,omitempty"`

	// 父目录id，没有则为根目录
	ParentId *int64 `json:"parent_id,omitempty"`

	// 上个节点ID,没有则为首节点
	PrevId *int64 `json:"prev_id,omitempty"`

	// 下个节点ID,没有则为尾节点
	NextId *int64 `json:"next_id,omitempty"`

	// 创建时传空，更新时必填
	Id *int64 `json:"id,omitempty"`

	// 认证ID，自动生成
	QualifiedId *string `json:"qualified_id,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 拥有业务指标数量，前端不传
	BizmetricNum *int32 `json:"bizmetric_num,omitempty"`

	// 拥有子流程的数量，不包括子流程的子流程
	ChildrenNum *int32 `json:"children_num,omitempty"`

	// 下层子目录
	Children       *[]BizCatalogVo `json:"children,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListCatalogTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogTreeResponse struct{}"
	}

	return strings.Join([]string{"ListCatalogTreeResponse", string(data)}, " ")
}
