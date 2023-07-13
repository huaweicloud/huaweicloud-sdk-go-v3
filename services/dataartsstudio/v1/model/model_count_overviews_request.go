package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountOverviewsRequest Request Object
type CountOverviewsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o CountOverviewsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountOverviewsRequest struct{}"
	}

	return strings.Join([]string{"CountOverviewsRequest", string(data)}, " ")
}
