package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 失败原因。如果为空，则集群处于正常状态。
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
