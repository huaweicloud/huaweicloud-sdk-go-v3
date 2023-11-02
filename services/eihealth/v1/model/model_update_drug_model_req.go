package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDrugModelReq struct {

	// 是否共享
	Shareable bool `json:"shareable"`
}

func (o UpdateDrugModelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDrugModelReq struct{}"
	}

	return strings.Join([]string{"UpdateDrugModelReq", string(data)}, " ")
}
