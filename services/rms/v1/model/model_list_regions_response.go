package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRegionsResponse struct {
	// 区域信息项列表

	Value          *[]Region `json:"value,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRegionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegionsResponse struct{}"
	}

	return strings.Join([]string{"ListRegionsResponse", string(data)}, " ")
}
