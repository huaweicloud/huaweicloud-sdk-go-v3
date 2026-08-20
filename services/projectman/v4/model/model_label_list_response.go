package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LabelListResponse struct {

	// 标签列表
	Data *[]LabelEntity `json:"data,omitempty"`

	// 标签总数
	Total *int32 `json:"total,omitempty"`
}

func (o LabelListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelListResponse struct{}"
	}

	return strings.Join([]string{"LabelListResponse", string(data)}, " ")
}
