package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeShareChargeModeV2Request Request Object
type ChangeShareChargeModeV2Request struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *ChangeShareChargeModeRequestBody `json:"body,omitempty"`
}

func (o ChangeShareChargeModeV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeShareChargeModeV2Request struct{}"
	}

	return strings.Join([]string{"ChangeShareChargeModeV2Request", string(data)}, " ")
}
