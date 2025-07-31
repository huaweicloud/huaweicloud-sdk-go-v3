package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalAssetScanTaskResponse Response Object
type CreateGlobalAssetScanTaskResponse struct {

	// 失败原因
	Error          *string `json:"error,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGlobalAssetScanTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalAssetScanTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateGlobalAssetScanTaskResponse", string(data)}, " ")
}
