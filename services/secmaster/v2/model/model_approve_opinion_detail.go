package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApproveOpinionDetail 审核详情信息
type ApproveOpinionDetail struct {

	// 审核结果
	Result *string `json:"result,omitempty"`

	// 审核内容
	Content *string `json:"content,omitempty"`
}

func (o ApproveOpinionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApproveOpinionDetail struct{}"
	}

	return strings.Join([]string{"ApproveOpinionDetail", string(data)}, " ")
}
