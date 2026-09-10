package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatasourceConnectionHostMessageRequest Request Object
type UpdateDatasourceConnectionHostMessageRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// DataArts Studio实例ID。
	InstanceId string `json:"instance_id"`

	// 连接ID，用于标识资源组网络连接的UUID。
	ConnectionId string `json:"connection_id"`

	Body *UpdateConnectionHostReq `json:"body,omitempty"`
}

func (o UpdateDatasourceConnectionHostMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatasourceConnectionHostMessageRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatasourceConnectionHostMessageRequest", string(data)}, " ")
}
