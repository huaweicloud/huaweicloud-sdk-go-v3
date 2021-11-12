package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 页码

	Page *int32 `json:"page,omitempty"`
	// 单页条数

	Pagesize *int32 `json:"pagesize,omitempty"`
	// 域名

	Hostname *string `json:"hostname,omitempty"`
	// 策略名

	Policyname *string `json:"policyname,omitempty"`
}

func (o ListHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostRequest struct{}"
	}

	return strings.Join([]string{"ListHostRequest", string(data)}, " ")
}
