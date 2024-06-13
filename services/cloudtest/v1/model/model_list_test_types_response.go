package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestTypesResponse Response Object
type ListTestTypesResponse struct {
	Value          *[]IntegerIdAndNameVo `json:"value,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListTestTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestTypesResponse struct{}"
	}

	return strings.Join([]string{"ListTestTypesResponse", string(data)}, " ")
}
