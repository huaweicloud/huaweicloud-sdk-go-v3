package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDisasterProgressResponse Response Object
type ShowDisasterProgressResponse struct {
	DisasterRecoveryProgress *ClusterDisasterRecovery `json:"disaster_recovery_progress,omitempty"`
	HttpStatusCode           int                      `json:"-"`
}

func (o ShowDisasterProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDisasterProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowDisasterProgressResponse", string(data)}, " ")
}
