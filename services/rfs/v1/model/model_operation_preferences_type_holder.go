package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OperationPreferencesTypeHolder struct {
	OperationPreferences *OperationPreferences `json:"operation_preferences,omitempty"`
}

func (o OperationPreferencesTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationPreferencesTypeHolder struct{}"
	}

	return strings.Join([]string{"OperationPreferencesTypeHolder", string(data)}, " ")
}
