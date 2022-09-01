package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListComponentInfosResponse struct {

	// 组件信息。
	Nodes *[]Nodes `json:"nodes,omitempty" xml:"nodes"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListComponentInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentInfosResponse struct{}"
	}

	return strings.Join([]string{"ListComponentInfosResponse", string(data)}, " ")
}
