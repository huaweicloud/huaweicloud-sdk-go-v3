package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatasetsResponse Response Object
type ListDatasetsResponse struct {

	// 数据集列表
	Datasets *[]Dataset `json:"datasets,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListDatasetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatasetsResponse struct{}"
	}

	return strings.Join([]string{"ListDatasetsResponse", string(data)}, " ")
}
