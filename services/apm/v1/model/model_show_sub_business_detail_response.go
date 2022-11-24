package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSubBusinessDetailResponse struct {

	// 子应用id。
	Id *int64 `json:"id,omitempty"`

	// 创建时间。
	GmtCreate *string `json:"gmt_create,omitempty"`

	// 修改时间。
	GmtModify *string `json:"gmt_modify,omitempty"`

	// 父亲的子应用id。
	ParentId *int64 `json:"parent_id,omitempty"`

	// 子应用的英文名称。
	Name *string `json:"name,omitempty"`

	// 子应用的展示名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 所属应用id。
	BusinessId *int64 `json:"business_id,omitempty"`

	// 内部租户id。
	InnerDomainId *int32 `json:"inner_domain_id,omitempty"`

	// 创建者的userId。
	CreatorId *int64 `json:"creator_id,omitempty"`

	// 应用的UUID。
	Uuid *string `json:"uuid,omitempty"`

	// 应用描述说明。
	Descp *string `json:"descp,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 修改时间。
	ModifyTime *string `json:"modify_time,omitempty"`

	// 创建者的用户名。
	CreatorName    *string `json:"creator_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSubBusinessDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubBusinessDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSubBusinessDetailResponse", string(data)}, " ")
}
