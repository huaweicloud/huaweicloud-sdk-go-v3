package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomersRegionsResponse Response Object
type ListCustomersRegionsResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 区域信息
	RegionInfos    *[]Region `json:"region_infos,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCustomersRegionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomersRegionsResponse struct{}"
	}

	return strings.Join([]string{"ListCustomersRegionsResponse", string(data)}, " ")
}
