package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigDelUserParams 获取DataArtsStudio工作空间角色信息
type ApigDelUserParams struct {

	// 用户组列表
	UserIds *interface{} `json:"user_ids"`
}

func (o ApigDelUserParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigDelUserParams struct{}"
	}

	return strings.Join([]string{"ApigDelUserParams", string(data)}, " ")
}
