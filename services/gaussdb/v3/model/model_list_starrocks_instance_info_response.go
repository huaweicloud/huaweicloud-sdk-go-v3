package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStarrocksInstanceInfoResponse Response Object
type ListStarrocksInstanceInfoResponse struct {

	// 实例信息。
	Instances      *[]StarRocksInstanceInfoInstances `json:"instances,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListStarrocksInstanceInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStarrocksInstanceInfoResponse struct{}"
	}

	return strings.Join([]string{"ListStarrocksInstanceInfoResponse", string(data)}, " ")
}
