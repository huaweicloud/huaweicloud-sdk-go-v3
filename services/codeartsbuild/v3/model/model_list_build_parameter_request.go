package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBuildParameterRequest Request Object
type ListBuildParameterRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 构建任务的构建编号，从1开始，每次构建递增1
	BuildNo int32 `json:"build_no"`
}

func (o ListBuildParameterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBuildParameterRequest struct{}"
	}

	return strings.Join([]string{"ListBuildParameterRequest", string(data)}, " ")
}
