package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDerivativeIndexesResponse Response Object
type ListDerivativeIndexesResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDerivativeIndexesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDerivativeIndexesResponse struct{}"
	}

	return strings.Join([]string{"ListDerivativeIndexesResponse", string(data)}, " ")
}
