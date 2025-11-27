package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsJobSpec struct {

	// 任务列表
	TaskList *[]UcsTask `json:"taskList,omitempty"`

	// 用户的domainID
	DomainID *string `json:"domainID,omitempty"`

	// 操作，create和retry二选一
	Operation *string `json:"operation,omitempty"`

	// Job等待时间，单位：秒
	WaitTimeOut *int32 `json:"waitTimeOut,omitempty"`

	// Job类别
	Type *string `json:"type,omitempty"`

	// 相关目标
	RelatedObjects map[string]string `json:"relatedObjects,omitempty"`

	// 额外参数
	ExtendParam map[string]interface{} `json:"extendParam,omitempty"`
}

func (o UcsJobSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsJobSpec struct{}"
	}

	return strings.Join([]string{"UcsJobSpec", string(data)}, " ")
}
