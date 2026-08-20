package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ThirdPartyAssociatedDto 查询外部链接结果
type ThirdPartyAssociatedDto struct {

	// 工作项归属项目的项目空间ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 新关联外部链接时会创建一条数据，该数据的唯一标识ID，可以在查询外部链接接口以及关联外部链接接口响应体中找到。
	Id *string `json:"id,omitempty"`

	// 工作项下关联外部链接的创建时间。
	CreatedDate *string `json:"created_date,omitempty"`

	// 工作项下关联外部链接的创建人。
	CreatedBy *string `json:"created_by,omitempty"`

	// 工作项下关联外部链接的名称。
	Title *string `json:"title,omitempty"`

	// 工作项下关联外部链接的地址。
	Url *string `json:"url,omitempty"`
}

func (o ThirdPartyAssociatedDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThirdPartyAssociatedDto struct{}"
	}

	return strings.Join([]string{"ThirdPartyAssociatedDto", string(data)}, " ")
}
