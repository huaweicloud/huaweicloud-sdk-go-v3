package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFlavorsResponse struct {

	// 规格列表。
	Flavors *[]Flavor `json:"flavors,omitempty" xml:"flavors"`

	PageInfo *PageInfo `json:"page_info,omitempty" xml:"page_info"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}
