package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWhetherUseCloudDbaResponse Response Object
type ShowWhetherUseCloudDbaResponse struct {

	// 能否使用云DBA功能
	CanUse         *bool `json:"can_use,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowWhetherUseCloudDbaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWhetherUseCloudDbaResponse struct{}"
	}

	return strings.Join([]string{"ShowWhetherUseCloudDbaResponse", string(data)}, " ")
}
