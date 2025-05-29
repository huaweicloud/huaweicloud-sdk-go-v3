package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobInfoRequest Request Object
type ShowJobInfoRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`
}

func (o ShowJobInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowJobInfoRequest", string(data)}, " ")
}
