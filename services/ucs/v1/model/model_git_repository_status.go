package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GitRepositoryStatus struct {

	// 控制器上次处理的GitRepository版本号
	ObservedGeneration *int32 `json:"observedGeneration,omitempty"`

	// GitRepository当前的条件集合，用于表示对象的不同状态
	Conditions *[]interface{} `json:"conditions,omitempty"`

	Artifact *Artifact `json:"artifact,omitempty"`
}

func (o GitRepositoryStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GitRepositoryStatus struct{}"
	}

	return strings.Join([]string{"GitRepositoryStatus", string(data)}, " ")
}
