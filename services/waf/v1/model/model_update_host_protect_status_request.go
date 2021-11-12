package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateHostProtectStatusRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 域名Id（通过查询云模式防护域名列表获取域名id)

	InstanceId string `json:"instance_id"`

	Body *UpdateHostProtectStatusRequestBody `json:"body,omitempty"`
}

func (o UpdateHostProtectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostProtectStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostProtectStatusRequest", string(data)}, " ")
}
