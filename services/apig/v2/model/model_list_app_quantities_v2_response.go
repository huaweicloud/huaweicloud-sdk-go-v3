package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppQuantitiesV2Response struct {
	// 已进行API访问授权的APP个数

	AuthedNums *int32 `json:"authed_nums,omitempty"`
	// 未进行API访问授权的APP个数

	UnauthedNums   *int32 `json:"unauthed_nums,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAppQuantitiesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppQuantitiesV2Response struct{}"
	}

	return strings.Join([]string{"ListAppQuantitiesV2Response", string(data)}, " ")
}
