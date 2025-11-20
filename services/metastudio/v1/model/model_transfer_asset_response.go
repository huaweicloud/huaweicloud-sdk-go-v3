package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferAssetResponse Response Object
type TransferAssetResponse struct {

	// 资产转移任务ID
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o TransferAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferAssetResponse struct{}"
	}

	return strings.Join([]string{"TransferAssetResponse", string(data)}, " ")
}
