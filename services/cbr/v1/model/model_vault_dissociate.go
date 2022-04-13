package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VaultDissociate struct {
	// 策略ID

	PolicyId string `json:"policy_id"`
}

func (o VaultDissociate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VaultDissociate struct{}"
	}

	return strings.Join([]string{"VaultDissociate", string(data)}, " ")
}
