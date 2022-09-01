package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRepositoryTagsResponse struct {

	// 镜像tag列表
	Body *[]ShowReposTagResp `json:"body,omitempty" xml:"body"`

	ContentRange   *string `json:"Content-Range,omitempty" xml:"Content-Range"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryTagsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryTagsResponse", string(data)}, " ")
}
