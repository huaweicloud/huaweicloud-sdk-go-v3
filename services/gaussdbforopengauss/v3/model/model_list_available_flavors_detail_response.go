package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableFlavorsDetailResponse Response Object
type ListAvailableFlavorsDetailResponse struct {
	FlavorCategory *FlavorCategoryDetail `json:"flavor_category,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListAvailableFlavorsDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableFlavorsDetailResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableFlavorsDetailResponse", string(data)}, " ")
}
