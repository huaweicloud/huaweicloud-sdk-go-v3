package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSpecsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 规格编码列表
	SpecCodes      *[]Spec `json:"spec_codes,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSpecsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecsResponse struct{}"
	}

	return strings.Join([]string{"ListSpecsResponse", string(data)}, " ")
}
