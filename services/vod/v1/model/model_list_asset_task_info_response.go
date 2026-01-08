package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetTaskInfoResponse Response Object
type ListAssetTaskInfoResponse struct {

	// 结果列表展示
	Results *[]Result `json:"results,omitempty"`

	// 根据条件的总条目数量
	Count *int64 `json:"count,omitempty"`

	// 下次查询的标志位
	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAssetTaskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetTaskInfoResponse struct{}"
	}

	return strings.Join([]string{"ListAssetTaskInfoResponse", string(data)}, " ")
}
