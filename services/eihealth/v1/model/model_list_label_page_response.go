package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLabelPageResponse Response Object
type ListLabelPageResponse struct {

	// 标签页面列表
	LabelPages *[]LabelPageListDto `json:"label_pages,omitempty"`

	// 标签页面总条数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLabelPageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelPageResponse struct{}"
	}

	return strings.Join([]string{"ListLabelPageResponse", string(data)}, " ")
}
