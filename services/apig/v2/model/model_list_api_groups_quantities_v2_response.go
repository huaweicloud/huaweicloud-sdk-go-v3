package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiGroupsQuantitiesV2Response Response Object
type ListApiGroupsQuantitiesV2Response struct {

	// 未上架的API分组个数  暂不支持
	OffsellNums *int32 `json:"offsell_nums,omitempty"`

	// 已上架的API分组个数
	OnsellNums     *int32 `json:"onsell_nums,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListApiGroupsQuantitiesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiGroupsQuantitiesV2Response struct{}"
	}

	return strings.Join([]string{"ListApiGroupsQuantitiesV2Response", string(data)}, " ")
}
