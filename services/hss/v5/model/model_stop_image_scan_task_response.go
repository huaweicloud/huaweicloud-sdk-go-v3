package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopImageScanTaskResponse Response Object
type StopImageScanTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopImageScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopImageScanTaskResponse struct{}"
	}

	return strings.Join([]string{"StopImageScanTaskResponse", string(data)}, " ")
}
