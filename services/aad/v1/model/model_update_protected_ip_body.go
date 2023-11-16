package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProtectedIpBody 更新防护ip tag请求体
type UpdateProtectedIpBody struct {

	// 防护ip的id
	Id string `json:"id"`

	// 本地标签
	Tag string `json:"tag"`
}

func (o UpdateProtectedIpBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectedIpBody struct{}"
	}

	return strings.Join([]string{"UpdateProtectedIpBody", string(data)}, " ")
}
