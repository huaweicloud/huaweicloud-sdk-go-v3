package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SourceRef 来源信息，Service信息或者模型信息
type SourceRef struct {

	// 一种资源ID，32~36位的英文、数字、短横组合
	Id string `json:"id"`

	// 一种资源ID，32~36位的英文、数字、短横组合
	VersionId *string `json:"version_id,omitempty"`
}

func (o SourceRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceRef struct{}"
	}

	return strings.Join([]string{"SourceRef", string(data)}, " ")
}
