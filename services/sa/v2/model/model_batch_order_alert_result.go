package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOrderAlertResult 委托告警返回对象
type BatchOrderAlertResult struct {

	// 失败id
	ErrorIds *[]string `json:"error_ids,omitempty"`

	// 成功败id
	SuccessIds *[]string `json:"success_ids,omitempty"`
}

func (o BatchOrderAlertResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOrderAlertResult struct{}"
	}

	return strings.Join([]string{"BatchOrderAlertResult", string(data)}, " ")
}
