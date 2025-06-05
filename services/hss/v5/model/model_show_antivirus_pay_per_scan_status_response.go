package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAntivirusPayPerScanStatusResponse Response Object
type ShowAntivirusPayPerScanStatusResponse struct {

	// 是否可用
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowAntivirusPayPerScanStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAntivirusPayPerScanStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowAntivirusPayPerScanStatusResponse", string(data)}, " ")
}
