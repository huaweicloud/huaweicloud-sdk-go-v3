package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeRegisteredRegionsResponse Response Object
type DescribeRegisteredRegionsResponse struct {

	// 局点列表
	Regions        *[]RegionDto `json:"regions,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DescribeRegisteredRegionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeRegisteredRegionsResponse struct{}"
	}

	return strings.Join([]string{"DescribeRegisteredRegionsResponse", string(data)}, " ")
}
