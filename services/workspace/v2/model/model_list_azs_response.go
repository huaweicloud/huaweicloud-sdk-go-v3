package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAzsResponse Response Object
type ListAzsResponse struct {

	// 支持类型。
	SupportType *[]string `json:"support_type,omitempty"`

	// 默认类型。
	DefaultType *string `json:"default_type,omitempty"`

	// 可用区。
	Azs            map[string][]AzInfo `json:"azs,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListAzsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAzsResponse struct{}"
	}

	return strings.Join([]string{"ListAzsResponse", string(data)}, " ")
}
