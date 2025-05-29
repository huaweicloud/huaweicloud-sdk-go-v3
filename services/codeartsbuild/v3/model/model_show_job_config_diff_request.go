package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobConfigDiffRequest Request Object
type ShowJobConfigDiffRequest struct {

	// 构建的任务ID； 编辑构建任务时，浏览器URL末尾的32位数字、字母组合的字符串。
	JobId string `json:"job_id"`

	// 新记录的序号
	RevisedlNo string `json:"revisedl_no"`

	// 原记录的序号
	OriginalNo string `json:"original_no"`
}

func (o ShowJobConfigDiffRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobConfigDiffRequest struct{}"
	}

	return strings.Join([]string{"ShowJobConfigDiffRequest", string(data)}, " ")
}
