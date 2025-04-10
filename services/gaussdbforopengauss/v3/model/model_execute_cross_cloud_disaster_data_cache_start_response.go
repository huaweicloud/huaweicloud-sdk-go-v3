package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteCrossCloudDisasterDataCacheStartResponse Response Object
type ExecuteCrossCloudDisasterDataCacheStartResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteCrossCloudDisasterDataCacheStartResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterDataCacheStartResponse struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterDataCacheStartResponse", string(data)}, " ")
}
