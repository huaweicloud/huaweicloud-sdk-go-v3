package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetbyReplicationInfoResponse Response Object
type CreateAssetbyReplicationInfoResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAssetbyReplicationInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetbyReplicationInfoResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetbyReplicationInfoResponse", string(data)}, " ")
}
