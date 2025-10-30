package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecreateVulScanTaskResponse Response Object
type RecreateVulScanTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RecreateVulScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecreateVulScanTaskResponse struct{}"
	}

	return strings.Join([]string{"RecreateVulScanTaskResponse", string(data)}, " ")
}
