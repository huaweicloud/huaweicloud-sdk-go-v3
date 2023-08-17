package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalUserExecuteInfo 用户用例执行信息
type ExternalUserExecuteInfo struct {
	Executor *NameAndIdVo `json:"executor,omitempty"`

	// 执行用例数
	ExecuteCount *int32 `json:"execute_count,omitempty"`
}

func (o ExternalUserExecuteInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalUserExecuteInfo struct{}"
	}

	return strings.Join([]string{"ExternalUserExecuteInfo", string(data)}, " ")
}
