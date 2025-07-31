package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMultiCloudClustersResponse Response Object
type UpdateMultiCloudClustersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMultiCloudClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMultiCloudClustersResponse struct{}"
	}

	return strings.Join([]string{"UpdateMultiCloudClustersResponse", string(data)}, " ")
}
