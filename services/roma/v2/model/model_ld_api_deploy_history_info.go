package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LdApiDeployHistoryInfo struct {

	// 部署的编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 部署的后端API编号
	LdApiId *string `json:"ld_api_id,omitempty" xml:"ld_api_id"`

	// 部署的前端API分组编号
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 部署的环境编号
	EnvId *string `json:"env_id,omitempty" xml:"env_id"`

	// 部署的前端API编号
	ApiId *string `json:"api_id,omitempty" xml:"api_id"`

	// 部署时间
	DeployTime *sdktime.SdkTime `json:"deploy_time,omitempty" xml:"deploy_time"`

	ApiDefinition *LdApiInfo `json:"api_definition,omitempty" xml:"api_definition"`
}

func (o LdApiDeployHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiDeployHistoryInfo struct{}"
	}

	return strings.Join([]string{"LdApiDeployHistoryInfo", string(data)}, " ")
}
