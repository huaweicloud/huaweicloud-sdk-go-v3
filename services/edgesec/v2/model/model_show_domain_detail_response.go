package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainDetailResponse Response Object
type ShowDomainDetailResponse struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`

	// 策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 防护状态： - 防护中：on - 未防护：off
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 创建域名的时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新域名的时间
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDomainDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainDetailResponse", string(data)}, " ")
}
