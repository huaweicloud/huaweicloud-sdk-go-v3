package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiffRefsParamDto struct {

	// 合并请求中原分支与目标分支的共同基准点
	BaseSha string `json:"base_sha"`

	// 合并请求中，从共同基准点开始到原分支方向的第一个提交点
	StartSha string `json:"start_sha"`

	// 合并请求中原分支指向的提交点
	HeadSha string `json:"head_sha"`
}

func (o DiffRefsParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiffRefsParamDto struct{}"
	}

	return strings.Join([]string{"DiffRefsParamDto", string(data)}, " ")
}
