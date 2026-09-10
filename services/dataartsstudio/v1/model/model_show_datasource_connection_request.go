package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatasourceConnectionRequest Request Object
type ShowDatasourceConnectionRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// DataArts Studio实例ID。
	InstanceId string `json:"instance_id"`

	// 连接ID，用于标识资源组网络连接的UUID。
	ConnectionId string `json:"connection_id"`
}

func (o ShowDatasourceConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatasourceConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowDatasourceConnectionRequest", string(data)}, " ")
}
