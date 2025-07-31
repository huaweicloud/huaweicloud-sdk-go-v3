package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAntivirusPayPerScanStatusResponse Response Object
type ChangeAntivirusPayPerScanStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeAntivirusPayPerScanStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAntivirusPayPerScanStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeAntivirusPayPerScanStatusResponse", string(data)}, " ")
}
