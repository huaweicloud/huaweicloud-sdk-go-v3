package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveMultiCloudClustersResponse Response Object
type RemoveMultiCloudClustersResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveMultiCloudClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveMultiCloudClustersResponse struct{}"
	}

	return strings.Join([]string{"RemoveMultiCloudClustersResponse", string(data)}, " ")
}
