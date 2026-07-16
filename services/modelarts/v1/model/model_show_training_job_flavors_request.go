package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingJobFlavorsRequest Request Object
type ShowTrainingJobFlavorsRequest struct {

	// 查询训练作业规格的类型，不填为查询所有。枚举值：  - CPU  - GPU  - [Ascend](tag:hc,hk,fcs_super)
	FlavorType *string `json:"flavor_type,omitempty"`
}

func (o ShowTrainingJobFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingJobFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ShowTrainingJobFlavorsRequest", string(data)}, " ")
}
