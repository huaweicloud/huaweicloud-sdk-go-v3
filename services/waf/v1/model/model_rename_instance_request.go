package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RenameInstanceRequest struct {

	// 通过企业项目管理服务的查询企业项目列表接口ListEnterpriseProject查询通过企业项目管理服务的查询企业项目列表接口ListEnterpriseProject查询企业项目ID
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
