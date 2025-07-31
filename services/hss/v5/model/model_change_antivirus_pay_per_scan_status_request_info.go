package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeAntivirusPayPerScanStatusRequestInfo struct {

	// 病毒查杀按次计费功能是否开启
	Enabled bool `json:"enabled"`
}

func (o ChangeAntivirusPayPerScanStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAntivirusPayPerScanStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeAntivirusPayPerScanStatusRequestInfo", string(data)}, " ")
}
