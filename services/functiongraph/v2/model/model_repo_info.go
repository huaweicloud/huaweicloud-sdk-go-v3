package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RepoInfo 存储库信息
type RepoInfo struct {

	// http地址
	HttpsUrl *string `json:"https_url,omitempty"`

	// 存储库链接
	WebUrl *string `json:"web_url,omitempty"`

	// 存储库状态
	RepoStatus *string `json:"repo_status,omitempty"`

	// 报错信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`
}

func (o RepoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoInfo struct{}"
	}

	return strings.Join([]string{"RepoInfo", string(data)}, " ")
}
