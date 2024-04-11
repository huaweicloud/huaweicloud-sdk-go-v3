package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCodeTableValuesResponse Response Object
type SearchCodeTableValuesResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SearchCodeTableValuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCodeTableValuesResponse struct{}"
	}

	return strings.Join([]string{"SearchCodeTableValuesResponse", string(data)}, " ")
}
