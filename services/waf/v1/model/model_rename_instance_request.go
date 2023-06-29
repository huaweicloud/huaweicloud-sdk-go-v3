package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RenameInstanceRequest Request Object
type RenameInstanceRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 独享引擎ID（通过调用WAF的ListInstance接口获取所有独享引擎信息查询独享引擎ID）
	InstanceId string `json:"instance_id"`

	Body *RenameInstanceRequestBody `json:"body,omitempty"`
}

func (o RenameInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameInstanceRequest struct{}"
	}

	return strings.Join([]string{"RenameInstanceRequest", string(data)}, " ")
}
