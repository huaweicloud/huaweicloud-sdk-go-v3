package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComputingResourceFlavorsRequest Request Object
type ListComputingResourceFlavorsRequest struct {

	// 可用区
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`
}

func (o ListComputingResourceFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComputingResourceFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListComputingResourceFlavorsRequest", string(data)}, " ")
}
