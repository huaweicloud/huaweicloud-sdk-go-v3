package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstanceResponse Response Object
type UpgradeInstanceResponse struct {

	// 提交升级任务id
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceResponse", string(data)}, " ")
}
