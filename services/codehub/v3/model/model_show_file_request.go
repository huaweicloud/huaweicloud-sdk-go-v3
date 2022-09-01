package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowFileRequest struct {

	// 仓库短id
	RepoId int32 `json:"repo_id" xml:"repo_id"`

	// Url编码的新文件的完整路径。
	FilePath string `json:"file_path" xml:"file_path"`

	// commit id，仓库的branch名或tag名
	Ref string `json:"ref" xml:"ref"`
}

func (o ShowFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileRequest struct{}"
	}

	return strings.Join([]string{"ShowFileRequest", string(data)}, " ")
}
