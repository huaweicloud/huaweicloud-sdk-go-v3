package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeCbhInstanceRequest Request Object
type UpgradeCbhInstanceRequest struct {
	Body *UpgradeCbhRequestBody `json:"body,omitempty"`
}

func (o UpgradeCbhInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeCbhInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpgradeCbhInstanceRequest", string(data)}, " ")
}
