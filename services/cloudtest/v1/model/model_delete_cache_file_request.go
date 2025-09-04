package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCacheFileRequest Request Object
type DeleteCacheFileRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 删除附件路径
	FilePath *string `json:"file_path,omitempty"`

	// 附件Uri
	Uri *string `json:"uri,omitempty"`

	// 附件挂载资源Uri
	ParentUri *string `json:"parent_uri,omitempty"`

	// 是否备份
	BakUp *bool `json:"bak_up,omitempty"`
}

func (o DeleteCacheFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCacheFileRequest struct{}"
	}

	return strings.Join([]string{"DeleteCacheFileRequest", string(data)}, " ")
}
