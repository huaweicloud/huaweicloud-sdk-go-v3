package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 失败原因。集群处于正常状态时不返回。
type FailedReasons struct {
	CreateFailed *FailedReasonsCreateFailed `json:"CREATE_FAILED,omitempty"`
}

func (o FailedReasons) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedReasons struct{}"
	}

	return strings.Join([]string{"FailedReasons", string(data)}, " ")
}
