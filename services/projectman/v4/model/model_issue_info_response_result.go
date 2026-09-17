package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueInfoResponseResult **参数解释：** 返回信息。
type IssueInfoResponseResult struct {
	Issue *IssueDetailResponseV2 `json:"issue,omitempty"`
}

func (o IssueInfoResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueInfoResponseResult struct{}"
	}

	return strings.Join([]string{"IssueInfoResponseResult", string(data)}, " ")
}
