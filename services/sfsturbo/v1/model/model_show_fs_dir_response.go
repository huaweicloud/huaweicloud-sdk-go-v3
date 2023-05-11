package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowFsDirResponse struct {

	// 目录全路径
	Path *string `json:"path,omitempty"`

	// 目录权限
	Mode *int64 `json:"mode,omitempty"`

	// 用户id
	Uid *int64 `json:"uid,omitempty"`

	// 用户组id
	Gid            *int64 `json:"gid,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowFsDirResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFsDirResponse struct{}"
	}

	return strings.Join([]string{"ShowFsDirResponse", string(data)}, " ")
}
