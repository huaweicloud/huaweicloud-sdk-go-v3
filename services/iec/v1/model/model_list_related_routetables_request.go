package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelatedRoutetablesRequest Request Object
type ListRelatedRoutetablesRequest struct {

	// 子网ID
	SubnetId string `json:"subnet_id"`
}

func (o ListRelatedRoutetablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelatedRoutetablesRequest struct{}"
	}

	return strings.Join([]string{"ListRelatedRoutetablesRequest", string(data)}, " ")
}
