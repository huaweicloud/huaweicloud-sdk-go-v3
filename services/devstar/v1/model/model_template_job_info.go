/*
 * DevStar
 *
 * DevStar API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type TemplateJobInfo struct {
	// 应用名称
	ApplicationName *string `json:"application_name,omitempty"`
	// 任务依赖的模板id
	TemplateId string `json:"template_id"`
	// 应用代码生成后的地址类型，目前支持0：codehub地址
	RepoType *int32 `json:"repo_type,omitempty"`
	// 应用的动态参数json
	Properties *interface{}    `json:"properties,omitempty"`
	RepoInfo   *RepositoryInfo `json:"repo_info,omitempty"`
}

func (o TemplateJobInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TemplateJobInfo", string(data)}, " ")
}
