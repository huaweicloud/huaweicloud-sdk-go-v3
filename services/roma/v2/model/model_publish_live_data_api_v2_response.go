package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishLiveDataApiV2Response Response Object
type PublishLiveDataApiV2Response struct {

	// 部署的编号
	Id *string `json:"id,omitempty"`

	// 部署的后端API编号
	LdApiId *string `json:"ld_api_id,omitempty"`

	// 部署的前端API分组编号
	GroupId *string `json:"group_id,omitempty"`

	// 部署的环境编号
	EnvId *string `json:"env_id,omitempty"`

	// 部署的前端API编号
	ApiId *string `json:"api_id,omitempty"`

	// 部署时间
	DeployTime *sdktime.SdkTime `json:"deploy_time,omitempty"`

	ApiDefinition  *LdApiInfo `json:"api_definition,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o PublishLiveDataApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishLiveDataApiV2Response struct{}"
	}

	return strings.Join([]string{"PublishLiveDataApiV2Response", string(data)}, " ")
}
