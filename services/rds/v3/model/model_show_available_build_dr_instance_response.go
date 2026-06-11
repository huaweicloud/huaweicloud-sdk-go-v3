package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableBuildDrInstanceResponse Response Object
type ShowAvailableBuildDrInstanceResponse struct {

	// 实例列表
	Instances      *[]DrInstance `json:"instances,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowAvailableBuildDrInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableBuildDrInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowAvailableBuildDrInstanceResponse", string(data)}, " ")
}
