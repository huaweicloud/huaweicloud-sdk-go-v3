package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepoDetailsResponse Response Object
type ListRepoDetailsResponse struct {

	// 镜像仓库列表.
	Repos *[]ShowReposRespV3 `json:"repos,omitempty"`

	// 分页查询时,查询下一页数据的起始位置。
	NextMarker     *int32 `json:"nextMarker,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRepoDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepoDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListRepoDetailsResponse", string(data)}, " ")
}
