package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupOffsitePolicyResponse Response Object
type UpdateBackupOffsitePolicyResponse struct {

	// 结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateBackupOffsitePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupOffsitePolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateBackupOffsitePolicyResponse", string(data)}, " ")
}
