package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryResourcesRequestBody struct {
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`

	Matches *[]Match `json:"matches,omitempty"`
}

func (o QueryResourcesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResourcesRequestBody struct{}"
	}

	return strings.Join([]string{"QueryResourcesRequestBody", string(data)}, " ")
}
