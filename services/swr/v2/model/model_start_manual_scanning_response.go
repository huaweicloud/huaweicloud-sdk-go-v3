package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartManualScanningResponse Response Object
type StartManualScanningResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartManualScanningResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartManualScanningResponse struct{}"
	}

	return strings.Join([]string{"StartManualScanningResponse", string(data)}, " ")
}
