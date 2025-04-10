package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteCrossCloudDisasterDataCacheEndResponse Response Object
type ExecuteCrossCloudDisasterDataCacheEndResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteCrossCloudDisasterDataCacheEndResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteCrossCloudDisasterDataCacheEndResponse struct{}"
	}

	return strings.Join([]string{"ExecuteCrossCloudDisasterDataCacheEndResponse", string(data)}, " ")
}
