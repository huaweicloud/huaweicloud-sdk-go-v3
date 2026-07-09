package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogRetentionCommonSettingsResponse Response Object
type UpdateLogRetentionCommonSettingsResponse struct {

	// 状态  - success：成功  - fail：失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLogRetentionCommonSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogRetentionCommonSettingsResponse struct{}"
	}

	return strings.Join([]string{"UpdateLogRetentionCommonSettingsResponse", string(data)}, " ")
}
