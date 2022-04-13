package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LdApiDeployHistoryInfo struct {
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

	ApiDefinition *LdApiInfo `json:"api_definition,omitempty"`
}

func (o LdApiDeployHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiDeployHistoryInfo struct{}"
	}

	return strings.Join([]string{"LdApiDeployHistoryInfo", string(data)}, " ")
}
