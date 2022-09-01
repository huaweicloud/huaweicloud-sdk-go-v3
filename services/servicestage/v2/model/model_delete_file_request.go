package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteFileRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth" xml:"X-Repo-Auth"`

	// 组织ID。
	Namespace string `json:"namespace" xml:"namespace"`

	// 仓库项目ID，如果含有“/”，需要将“/”替换为“:”。
	Project string `json:"project" xml:"project"`

	// 文件路径，需要将“/”替换为“:”。
	Path string `json:"path" xml:"path"`

	// 分支名称或者tag标签名称或者commit sha。
	Ref string `json:"ref" xml:"ref"`

	// 提交描述。
	Message string `json:"message" xml:"message"`

	// 最后一次提交的commit sha值。
	Sha string `json:"sha" xml:"sha"`
}

func (o DeleteFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFileRequest struct{}"
	}

	return strings.Join([]string{"DeleteFileRequest", string(data)}, " ")
}
