package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PublisherVo struct {

	// 代码地址
	CodeRepo *string `json:"code_repo,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 发布商描述
	Description *string `json:"description,omitempty"`

	// EAMAP注册信息
	EamapInfo *string `json:"eamap_info,omitempty"`

	// 邮箱
	Email *string `json:"email,omitempty"`

	// espase交流群号
	EspaseGroup *string `json:"espase_group,omitempty"`

	// 插件数量
	ExtensionCount *int64 `json:"extension_count,omitempty"`

	// 发布商ID
	Id *string `json:"id,omitempty"`

	// 唯一标志
	Identifier *string `json:"identifier,omitempty"`

	// 开源发布商,0:非开源; 1:开源;
	IsOpen *bool `json:"is_open,omitempty"`

	// 发布商或组织,0:发布商; 1:组织;
	IsOrg *bool `json:"is_org,omitempty"`

	// 发布商logo
	LogoUrl *string `json:"logo_url,omitempty"`

	// 成员数量
	MemberCount *int64 `json:"member_count,omitempty"`

	// 发布商名称
	Name *string `json:"name,omitempty"`

	// 是否是官方发布商
	Official *bool `json:"official,omitempty"`

	// 成员角色
	Owners *[]MemberRoleVo `json:"owners,omitempty"`

	// 是否开启发布商审核,1:开启；0:关闭
	PublisherReview *bool `json:"publisher_review,omitempty"`

	// 角色
	Role *string `json:"role,omitempty"`

	// 状态,0:禁用; 1:正常;
	Status *bool `json:"status,omitempty"`

	// 匹配数量
	SuiteCount *int64 `json:"suite_count,omitempty"`

	// 支持地址
	SupportUrl *string `json:"support_url,omitempty"`

	// 是否忽略系统审核,1:忽略；0:不忽略
	SystemReview *bool `json:"system_review,omitempty"`

	// 更新时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`

	// 官网地址
	WebUrl *string `json:"web_url,omitempty"`
}

func (o PublisherVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublisherVo struct{}"
	}

	return strings.Join([]string{"PublisherVo", string(data)}, " ")
}
