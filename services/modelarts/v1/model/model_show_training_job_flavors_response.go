package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingJobFlavorsResponse Response Object
type ShowTrainingJobFlavorsResponse struct {

	// 训练作业资源规格总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 训练作业资源规格列表。
	Flavors        *[]FlavorResponseWithSupport `json:"flavors,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowTrainingJobFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingJobFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ShowTrainingJobFlavorsResponse", string(data)}, " ")
}
