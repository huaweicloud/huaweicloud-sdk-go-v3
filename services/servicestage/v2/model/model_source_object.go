package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SourceObject 组件来源。
type SourceObject struct {
	Kind *SourceKind `json:"kind"`

	Spec *SourceOrArtifact `json:"spec"`
}

func (o SourceObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceObject struct{}"
	}

	return strings.Join([]string{"SourceObject", string(data)}, " ")
}
