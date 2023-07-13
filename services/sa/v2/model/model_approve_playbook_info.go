package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApprovePlaybookInfo approve of playbook
type ApprovePlaybookInfo struct {

	// 审核结果  通过：PASS 不通过：UN_PASS
	Result *string `json:"result,omitempty"`

	// 审核意见
	Content *string `json:"content,omitempty"`
}

func (o ApprovePlaybookInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovePlaybookInfo struct{}"
	}

	return strings.Join([]string{"ApprovePlaybookInfo", string(data)}, " ")
}
