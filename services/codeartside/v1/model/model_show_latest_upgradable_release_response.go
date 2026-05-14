package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLatestUpgradableReleaseResponse Response Object
type ShowLatestUpgradableReleaseResponse struct {
	Result *ShowLatestUpgradableReleaseResponseBodyResult `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLatestUpgradableReleaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestUpgradableReleaseResponse struct{}"
	}

	return strings.Join([]string{"ShowLatestUpgradableReleaseResponse", string(data)}, " ")
}
