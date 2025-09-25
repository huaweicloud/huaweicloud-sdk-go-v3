package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManualImageScanTaskResponse Response Object
type CreateManualImageScanTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateManualImageScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualImageScanTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateManualImageScanTaskResponse", string(data)}, " ")
}
