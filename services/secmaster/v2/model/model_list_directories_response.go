package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDirectoriesResponse Response Object
type ListDirectoriesResponse struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 目录列表
	Directories *[]string `json:"directories,omitempty"`

	// 目录I18N列表
	DirectoryI18ns *[]DirectoryI18N `json:"directory_i18ns,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListDirectoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDirectoriesResponse struct{}"
	}

	return strings.Join([]string{"ListDirectoriesResponse", string(data)}, " ")
}
