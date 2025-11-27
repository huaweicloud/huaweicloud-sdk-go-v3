package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SourceRef struct {

	// 被引用资源的类型，必须为 OCIRepository、GitRepository或Bucket
	Kind string `json:"kind"`

	// 被引用资源的名称
	Name string `json:"name"`
}

func (o SourceRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceRef struct{}"
	}

	return strings.Join([]string{"SourceRef", string(data)}, " ")
}
