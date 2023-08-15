package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLabelsResponse Response Object
type ListLabelsResponse struct {

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 标签列表
	LabelList      *[]CaseLabelInfo `json:"label_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListLabelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelsResponse struct{}"
	}

	return strings.Join([]string{"ListLabelsResponse", string(data)}, " ")
}
