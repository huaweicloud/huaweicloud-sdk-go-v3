package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAvailableFlavorsResponse struct {

	// 实例可变更规格信息列表。
	Flavors *[]AvailableFlavorInfoResult `json:"flavors,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAvailableFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableFlavorsResponse", string(data)}, " ")
}
