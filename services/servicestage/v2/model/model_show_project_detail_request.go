package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowProjectDetailRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth"`

	// 仓库克隆URL。
	CloneUrl string `json:"clone_url"`
}

func (o ShowProjectDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectDetailRequest", string(data)}, " ")
}
