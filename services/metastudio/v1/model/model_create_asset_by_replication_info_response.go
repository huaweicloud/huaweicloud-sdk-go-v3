package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetByReplicationInfoResponse Response Object
type CreateAssetByReplicationInfoResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAssetByReplicationInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetByReplicationInfoResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetByReplicationInfoResponse", string(data)}, " ")
}
