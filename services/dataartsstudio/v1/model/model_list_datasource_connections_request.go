package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatasourceConnectionsRequest Request Object
type ListDatasourceConnectionsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// DataArts Studio实例ID。
	InstanceId string `json:"instance_id"`

	// 查询最大连接个数，默认100。
	Limit *int32 `json:"limit,omitempty"`

	// 连接名称，精确匹配。
	Name *string `json:"name,omitempty"`

	// 查询结果偏移量，默认为0（连接以创建时间进行排序）。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListDatasourceConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasourceConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListDatasourceConnectionsRequest", string(data)}, " ")
}
