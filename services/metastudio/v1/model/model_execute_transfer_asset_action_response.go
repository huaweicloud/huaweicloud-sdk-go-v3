package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTransferAssetActionResponse Response Object
type ExecuteTransferAssetActionResponse struct {
	AcceptRespone *TransAcceptResponse `json:"accept_respone,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteTransferAssetActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTransferAssetActionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteTransferAssetActionResponse", string(data)}, " ")
}
