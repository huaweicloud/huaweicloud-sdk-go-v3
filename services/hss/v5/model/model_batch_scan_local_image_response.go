package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchScanLocalImageResponse Response Object
type BatchScanLocalImageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchScanLocalImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchScanLocalImageResponse struct{}"
	}

	return strings.Join([]string{"BatchScanLocalImageResponse", string(data)}, " ")
}
