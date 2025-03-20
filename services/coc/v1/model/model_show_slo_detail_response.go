package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSloDetailResponse Response Object
type ShowSloDetailResponse struct {

	// SLO的ID
	Id *string `json:"id,omitempty"`

	// 应用名称
	ApplicationName *string `json:"application_name,omitempty"`

	// 应用ID
	ApplicationId *string `json:"application_id,omitempty"`

	// SLO的目标值
	SloTargets *float64 `json:"slo_targets,omitempty"`

	// SLi列表
	SliList        *[]SliDetail `json:"sli_list,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowSloDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSloDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSloDetailResponse", string(data)}, " ")
}
