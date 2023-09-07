package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIncrementComponentsDetailResponse Response Object
type ShowIncrementComponentsDetailResponse struct {

	// 更新时间。
	UpdateTime *string `json:"update_time,omitempty"`

	// 增量组件详情。
	IncrementComponentsList *[]IncreComponentDetail `json:"increment_components_list,omitempty"`
	HttpStatusCode          int                     `json:"-"`
}

func (o ShowIncrementComponentsDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIncrementComponentsDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowIncrementComponentsDetailResponse", string(data)}, " ")
}
