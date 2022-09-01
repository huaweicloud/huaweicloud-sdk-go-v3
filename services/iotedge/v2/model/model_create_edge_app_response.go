package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEdgeAppResponse struct {

	// 应用名称
	EdgeAppId *string `json:"edge_app_id,omitempty" xml:"edge_app_id"`

	// 应用描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// 最新发布版本
	LastPublishedVersion *string `json:"last_published_version,omitempty" xml:"last_published_version"`

	// 应用类型SYSTEM_REQUIRED|SYSTEM_OPTIONAL|USER
	AppType *string `json:"app_type,omitempty" xml:"app_type"`

	// 应用类型DATA_PROCESSING|PROTOCOL_PARSING
	FunctionType *string `json:"function_type,omitempty" xml:"function_type"`

	// 部署类型docker|process
	DeployType     *string `json:"deploy_type,omitempty" xml:"deploy_type"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEdgeAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeAppResponse struct{}"
	}

	return strings.Join([]string{"CreateEdgeAppResponse", string(data)}, " ")
}
