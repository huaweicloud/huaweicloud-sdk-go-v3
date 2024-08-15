package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MenusRsp struct {

	// 菜单ID。
	MenuId *string `json:"menu_id,omitempty"`

	// 服务号ID。
	PubId *string `json:"pub_id,omitempty"`

	// 服务号名称。
	PubName *string `json:"pub_name,omitempty"`

	// LOGO图片资源ID。
	LogoImg *string `json:"logo_img,omitempty"`

	// LOGO图片资源URL。
	LogoUrl *string `json:"logo_url,omitempty"`

	Menu *Menus `json:"menu,omitempty"`

	// 资源状态。  - 1：未生效 - 2：已生效 - 3：已失效  - 4：已冻结
	MenuState *int32 `json:"menu_state,omitempty"`

	// 审核状态。 - 1：待审核  - 2：通过  - 3：驳回
	ApproveState *int32 `json:"approve_state,omitempty"`

	// 上线时间。格式为：2020-12-12T12:00:00Z。
	OnlineTime *string `json:"online_time,omitempty"`

	// 最新操作时间。格式为：2020-12-12T12:00:00Z。
	OperTime *string `json:"oper_time,omitempty"`

	// 企业ID。
	CompanyId *string `json:"company_id,omitempty"`

	// 企业名称。
	CompanyName *string `json:"company_name,omitempty"`
}

func (o MenusRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MenusRsp struct{}"
	}

	return strings.Join([]string{"MenusRsp", string(data)}, " ")
}
