package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFsDirRequestBody 创建目录请求
type CreateFsDirRequestBody struct {

	// 合法的的目录全路径
	Path string `json:"path"`

	// 目录权限
	Mode *int64 `json:"mode,omitempty"`

	// 用户id
	Uid *int64 `json:"uid,omitempty"`

	// 用户组id
	Gid *int64 `json:"gid,omitempty"`
}

func (o CreateFsDirRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFsDirRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFsDirRequestBody", string(data)}, " ")
}
