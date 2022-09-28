package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLabelResponse struct {

	// 标签列表个数
	Count *int64 `json:"count,omitempty"`

	// 标签详情
	Labels         *[]LabelRsp `json:"labels,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelResponse struct{}"
	}

	return strings.Join([]string{"ListLabelResponse", string(data)}, " ")
}
