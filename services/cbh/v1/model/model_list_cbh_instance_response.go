package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCbhInstanceResponse struct {
	// 实例总数

	Total *int32 `json:"total,omitempty"`

	QuotaDetail *QuotaDetail `json:"quota_detail,omitempty"`
	// 实例列表

	Instance       *[]InstanceDetail `json:"instance,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListCbhInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCbhInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListCbhInstanceResponse", string(data)}, " ")
}
