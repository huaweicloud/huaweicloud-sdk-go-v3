package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvTagsResponse Response Object
type ListEnvTagsResponse struct {

	// 环境标签数据列表。
	EnvTags *[]CmdbTagEntity `json:"env_tags,omitempty"`

	// 总条数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnvTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvTagsResponse struct{}"
	}

	return strings.Join([]string{"ListEnvTagsResponse", string(data)}, " ")
}
