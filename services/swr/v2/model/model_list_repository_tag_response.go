package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryTagResponse Response Object
type ListRepositoryTagResponse struct {

	// 镜像tag列表.
	Tags *[]ShowReposTagRespV3 `json:"tags,omitempty"`

	// 下次分页查询时的起始位置。
	NextMarker *string `json:"nextMarker,omitempty"`

	// 表示分页查询时是否还有下一页
	HasMore        *bool `json:"has_more,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListRepositoryTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryTagResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryTagResponse", string(data)}, " ")
}
