package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterRegionReqBody the request body of RegisterRegion
type RegisterRegionReqBody struct {

	// 局点ID
	RegionId string `json:"region_id"`
}

func (o RegisterRegionReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterRegionReqBody struct{}"
	}

	return strings.Join([]string{"RegisterRegionReqBody", string(data)}, " ")
}
