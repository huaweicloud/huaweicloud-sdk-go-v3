package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateHostRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 域名id，您可以通过调用查询云模式防护域名列表（ListHost）获取域名id

	InstanceId string `json:"instance_id"`

	Body *UpdateHostRequestBody `json:"body,omitempty"`
}

func (o UpdateHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostRequest", string(data)}, " ")
}
