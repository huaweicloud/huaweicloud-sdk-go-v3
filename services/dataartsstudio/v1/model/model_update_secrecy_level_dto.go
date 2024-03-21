package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecrecyLevelDto struct {

	// 密级描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateSecrecyLevelDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecrecyLevelDto struct{}"
	}

	return strings.Join([]string{"UpdateSecrecyLevelDto", string(data)}, " ")
}
