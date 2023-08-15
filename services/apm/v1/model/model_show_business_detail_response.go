package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBusinessDetailResponse Response Object
type ShowBusinessDetailResponse struct {

	// 应用id。
	Id *int64 `json:"id,omitempty"`

	// 创建时间。
	GmtCreate *string `json:"gmt_create,omitempty"`

	// 修改时间。
	GmtModify *string `json:"gmt_modify,omitempty"`

	// 是否是默认的应用。
	Default *bool `json:"default,omitempty"`

	// 应用的英文名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 应用的展示名称。
	Name *string `json:"name,omitempty"`

	// 是否是默认的应用。
	IsDefault *bool `json:"is_default,omitempty"`

	// 内部租户id。
	InnerDomainId *int32 `json:"inner_domain_id,omitempty"`

	// 企业项目的id。
	EpsId *string `json:"eps_id,omitempty"`

	// 创建者的userId。
	CreatorId *int64 `json:"creator_id,omitempty"`

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

func (o ShowBusinessDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBusinessDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowBusinessDetailResponse", string(data)}, " ")
}
