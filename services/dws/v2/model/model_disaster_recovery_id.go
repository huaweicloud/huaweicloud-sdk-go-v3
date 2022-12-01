package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisasterRecoveryId struct {

	// 容灾ID
	Id *string `json:"id,omitempty"`
}

func (o DisasterRecoveryId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecoveryId struct{}"
	}

	return strings.Join([]string{"DisasterRecoveryId", string(data)}, " ")
}
