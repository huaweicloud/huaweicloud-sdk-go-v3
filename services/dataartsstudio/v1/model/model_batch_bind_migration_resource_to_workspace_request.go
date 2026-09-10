package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBindMigrationResourceToWorkspaceRequest Request Object
type BatchBindMigrationResourceToWorkspaceRequest struct {

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// DataArts Studio实例ID。
	InstanceId string `json:"instance_id"`

	Body *BatchBindMigrationResourceToWorkspaceRequestBody `json:"body,omitempty"`
}

func (o BatchBindMigrationResourceToWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBindMigrationResourceToWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"BatchBindMigrationResourceToWorkspaceRequest", string(data)}, " ")
}
