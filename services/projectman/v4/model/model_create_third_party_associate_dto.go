package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateThirdPartyAssociateDto 关联外部链接返回结果
type CreateThirdPartyAssociateDto struct {

	// 租户唯一标识ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 工作项下关联外部链接的名称。
	Title *string `json:"title,omitempty"`

	// 外部链接的类别。
	Type *string `json:"type,omitempty"`

	// 工作项下关联外部链接的修改时间。
	ModifiedDate *string `json:"modified_date,omitempty"`

	// 工作项下关联外部链接的创建人。
	CreatedBy *string `json:"created_by,omitempty"`

	// 工作项下关联外部链接的地址。
	Url *string `json:"url,omitempty"`

	// 租户下项目唯一标识ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 工作项实例对应的唯一标识ID。
	WorkitemId *string `json:"workitem_id,omitempty"`

	// 工作项下关联外部链接的修改人。
	ModifiedBy *string `json:"modified_by,omitempty"`

	// 外部链接操作项ID。
	OperationId *string `json:"operation_id,omitempty"`

	// 新关联外部链接时会创建一条数据，该数据的唯一标识ID，可以在查询外部链接接口以及关联外部链接接口响应体中找到。
	Id *string `json:"id,omitempty"`

	// 工作项下关联外部链接的创建时间。
	CreatedDate *string `json:"created_date,omitempty"`

	// 外部链接的生命周期。
	State *string `json:"state,omitempty"`

	// 外部链接的类型。
	Category *string `json:"category,omitempty"`

	// 区域 。
	Region *string `json:"region,omitempty"`
}

func (o CreateThirdPartyAssociateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateThirdPartyAssociateDto struct{}"
	}

	return strings.Join([]string{"CreateThirdPartyAssociateDto", string(data)}, " ")
}
