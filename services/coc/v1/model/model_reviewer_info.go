package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReviewerInfo 审批人信息
type ReviewerInfo struct {

	// 审批人名称（IAM用户名）
	ReviewerName string `json:"reviewer_name"`

	// 审批人ID（IAM用户Id）
	ReviewerId string `json:"reviewer_id"`
}

func (o ReviewerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewerInfo struct{}"
	}

	return strings.Join([]string{"ReviewerInfo", string(data)}, " ")
}
