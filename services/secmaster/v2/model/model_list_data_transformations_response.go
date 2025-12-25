package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataTransformationsResponse Response Object
type ListDataTransformationsResponse struct {

	// 数量
	Count *int64 `json:"count,omitempty"`

	// 记录
	Records        *[]DataTransformationItem `json:"records,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListDataTransformationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataTransformationsResponse struct{}"
	}

	return strings.Join([]string{"ListDataTransformationsResponse", string(data)}, " ")
}
