package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateFileRequest struct {
	// 授权名称。

	XRepoAuth string `json:"X-Repo-Auth"`
	// 组织ID。

	Namespace string `json:"namespace"`
	// 仓库项目ID，如果含有“/”，需要将“/”替换为“:”。

	Project string `json:"project"`
	// 文件路径，需要将“/”替换为“:”。

	Path string `json:"path"`
	// 分支名称或者tag标签名称或者commit sha。

	Ref string `json:"ref"`

	Body *FileCreate `json:"body,omitempty"`
}

func (o CreateFileRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateFileRequest struct{}"
	}

	return strings.Join([]string{"CreateFileRequest", string(data)}, " ")
}
