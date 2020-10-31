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

type CodehubJobInfo struct {
	// 应用名称
	ApplicationName string `json:"application_name"`
	// 仓库是否私有
	Privately *bool `json:"privately,omitempty"`
	// 仓库短id
	ShortId *string `json:"short_id,omitempty"`
	// 代码存放的ssh地址
	CodeUrl *string `json:"code_url,omitempty"`
	// 区域id
	RegionId *string `json:"region_id,omitempty"`
	// 应用代码生成后的地址类型，目前支持0：codehub地址; 1：压缩包下载地址
	RepoType *int32 `json:"repo_type,omitempty"`
	// 应用的动态参数json
	Properties *interface{}    `json:"properties,omitempty"`
	RepoInfo   *RepositoryInfo `json:"repo_info,omitempty"`
}

func (o CodehubJobInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CodehubJobInfo", string(data)}, " ")
}
