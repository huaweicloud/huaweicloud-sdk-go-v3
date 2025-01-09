package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAutoInstallAppsReq 自动安装应用请求。 assign_scope=ASSIGN_USER: users必须非空，如果用户不满足对应应用的访问权限，会自动添加对应的权限。 assign_scope=ALL_USER: 会同时修改所有应用授权为全部用户。
type BatchAutoInstallAppsReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 50]。
	AppIds []string `json:"app_ids"`

	AssignScope *AssignType `json:"assign_scope"`

	// 用户列表，一次请求数量区间 [1, 50]。
	Users *[]AccountInfo `json:"users,omitempty"`
}

func (o BatchAutoInstallAppsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAutoInstallAppsReq struct{}"
	}

	return strings.Join([]string{"BatchAutoInstallAppsReq", string(data)}, " ")
}
