package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateHostProtectStatusRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 域名id，您可以通过调用查询云模式防护域名列表（ListHost）获取域名id
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *UpdateHostProtectStatusRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateHostProtectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostProtectStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateHostProtectStatusRequest", string(data)}, " ")
}
