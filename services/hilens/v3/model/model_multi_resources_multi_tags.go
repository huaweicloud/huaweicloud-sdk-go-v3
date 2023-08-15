package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultiResourcesMultiTags struct {

	// 资源列表
	Resources *[]NodeResource `json:"resources,omitempty"`

	// 标签列表
	Tags *[]NodeTag `json:"tags,omitempty"`
}

func (o MultiResourcesMultiTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiResourcesMultiTags struct{}"
	}

	return strings.Join([]string{"MultiResourcesMultiTags", string(data)}, " ")
}
