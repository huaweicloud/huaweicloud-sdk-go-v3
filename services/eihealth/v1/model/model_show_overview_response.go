package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOverviewResponse Response Object
type ShowOverviewResponse struct {

	// 平台ID
	Id *string `json:"id,omitempty"`

	// 存储量，单位：byte
	Storage *int64 `json:"storage,omitempty"`

	// 项目总数
	ProjectNum *int32 `json:"project_num,omitempty"`

	// 计费模式
	ChargeMode *int32 `json:"charge_mode,omitempty"`

	// 是否欠费
	IsArrears      *bool `json:"is_arrears,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowOverviewResponse", string(data)}, " ")
}
