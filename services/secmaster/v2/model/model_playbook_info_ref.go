package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookInfoRef 剧本信息
type PlaybookInfoRef struct {

	// 剧本ID
	Id *string `json:"id,omitempty"`

	// 剧本版本ID
	VersionId *string `json:"version_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`
}

func (o PlaybookInfoRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookInfoRef struct{}"
	}

	return strings.Join([]string{"PlaybookInfoRef", string(data)}, " ")
}
