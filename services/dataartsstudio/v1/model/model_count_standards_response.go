package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountStandardsResponse Response Object
type CountStandardsResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CountStandardsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountStandardsResponse struct{}"
	}

	return strings.Join([]string{"CountStandardsResponse", string(data)}, " ")
}
