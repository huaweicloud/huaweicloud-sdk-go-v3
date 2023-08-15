package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDimensionsResponse Response Object
type ListDimensionsResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDimensionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionsResponse struct{}"
	}

	return strings.Join([]string{"ListDimensionsResponse", string(data)}, " ")
}
