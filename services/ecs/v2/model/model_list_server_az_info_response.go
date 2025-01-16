package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerAzInfoResponse Response Object
type ListServerAzInfoResponse struct {

	// az详情信息
	AvailabilityZones *[]ListServerAzInfo `json:"availability_zones,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ListServerAzInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerAzInfoResponse struct{}"
	}

	return strings.Join([]string{"ListServerAzInfoResponse", string(data)}, " ")
}
