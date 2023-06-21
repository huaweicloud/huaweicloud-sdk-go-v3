package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 主页模型。
type PortalModel struct {

	// 主页ID。
	PortalId string `json:"portal_id"`

	// 企业ID。
	CompanyId *string `json:"company_id,omitempty"`

	// 企业名称。
	CompanyName string `json:"company_name"`

	// 服务号ID。
	PubId string `json:"pub_id"`

	// 服务号名称。
	PubName string `json:"pub_name"`

	// 主页背景图片资源ID。
	BackgroundImg string `json:"background_img"`

	// 背景图片URL。
	BackgroundImgUrl *string `json:"background_img_url,omitempty"`

	// 简介。
	Summary string `json:"summary"`

	// 热线号列表。  > 以JSON列表返回，格式： > {\"tel\": \"400-800-8800\", \"usage\": \"官方服务电话\"}。
	Tels string `json:"tels"`

	// 快应用列表。  > 以JSON列表返回，格式： > {\"name\": \"快应用名称\",\"logo_img\": \"快应用LOGO图片资源ID\", \"logo_img_url\": \"快应用LOGO图片资源URL\", \"description\": \"快应用描述\",\"deeplink\": \"hap://app/fastapp\",\"depend_engine_version\": \"1040\"}。
	Fastapps string `json:"fastapps"`

	// 资源状态。  - 1：未生效  - 2：已生效  - 3：已失效  - 4：已冻结
	State int32 `json:"state"`

	// 审核状态。  - 1：待审核  - 2：通过  - 3：驳回
	ApproveState int32 `json:"approve_state"`

	// 上线时间。
	OnlineTime *sdktime.SdkTime `json:"online_time"`

	// 创建人。
	Creator string `json:"creator"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 修改原因。
	ChangeReason string `json:"change_reason"`
}

func (o PortalModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortalModel struct{}"
	}

	return strings.Join([]string{"PortalModel", string(data)}, " ")
}
