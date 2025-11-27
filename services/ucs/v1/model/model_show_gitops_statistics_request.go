package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGitopsStatisticsRequest Request Object
type ShowGitopsStatisticsRequest struct {

	// 分页获取列表时，页的大小，默认为-1
	Limit *int32 `json:"limit,omitempty"`

	// 分页获取列表时，起始偏移量，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 分页获取列表时，排序参数，支持create_at和update_at，默认create_at： - create_at：按创建时间排序 - update_at：按更新时间排序
	OrderBy *string `json:"order_by,omitempty"`

	// 分页获取列表时，排序方向，支持desc和asc，默认desc： - desc：降序排序 - asc：升序排序
	Order *string `json:"order,omitempty"`
}

func (o ShowGitopsStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGitopsStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowGitopsStatisticsRequest", string(data)}, " ")
}
