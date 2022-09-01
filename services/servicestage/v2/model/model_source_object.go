package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 组件来源。
type SourceObject struct {
	Kind *SourceKind `json:"kind,omitempty" xml:"kind"`

	Spec *SourceOrArtifact `json:"spec,omitempty" xml:"spec"`
}

func (o SourceObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceObject struct{}"
	}

	return strings.Join([]string{"SourceObject", string(data)}, " ")
}
