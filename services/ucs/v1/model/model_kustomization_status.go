package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KustomizationStatus struct {

	// 最近一次成功协调的资源版本号，用于标识控制器已处理的对象代
	ObservedGeneration *int32 `json:"observedGeneration,omitempty"`

	// 当前对象的状态条件列表
	Conditions *[]interface{} `json:"conditions,omitempty"`

	// 最近一次成功应用的版本号
	LastAttemptedRevision *string `json:"lastAttemptedRevision,omitempty"`
}

func (o KustomizationStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KustomizationStatus struct{}"
	}

	return strings.Join([]string{"KustomizationStatus", string(data)}, " ")
}
