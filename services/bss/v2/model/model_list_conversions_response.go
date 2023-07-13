package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConversionsResponse Response Object
type ListConversionsResponse struct {

	// 度量单位的换算信息，具体参见表3。
	Conversions    *[]Conversion `json:"conversions,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListConversionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConversionsResponse struct{}"
	}

	return strings.Join([]string{"ListConversionsResponse", string(data)}, " ")
}
