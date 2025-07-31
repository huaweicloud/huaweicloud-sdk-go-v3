package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaselineDirectoryResponse Response Object
type ShowBaselineDirectoryResponse struct {
	TaskCondition *SecurityCheckTaskCondition `json:"task_condition,omitempty"`

	// **参数解释** 基线检查策略目录
	BaselineDirectoryList *[]ShowBaselineDirectoryInfo `json:"baseline_directory_list,omitempty"`

	// **参数解释** 基线检查策略目录
	PwdDirectoryList *[]ShowPwdDirectoryInfo `json:"pwd_directory_list,omitempty"`
	HttpStatusCode   int                     `json:"-"`
}

func (o ShowBaselineDirectoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaselineDirectoryResponse struct{}"
	}

	return strings.Join([]string{"ShowBaselineDirectoryResponse", string(data)}, " ")
}
