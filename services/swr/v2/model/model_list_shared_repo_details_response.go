package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSharedRepoDetailsResponse Response Object
type ListSharedRepoDetailsResponse struct {

	// 镜像仓库列表.
	Repos *[]ShowReposRespV3 `json:"repos,omitempty"`

	// 分页查询时,查询下一页数据的起始位置。
	NextMarker     *int32 `json:"nextMarker,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSharedRepoDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharedRepoDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSharedRepoDetailsResponse", string(data)}, " ")
}
