package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityDlfDataWareHousesRequest Request Object
type ListSecurityDlfDataWareHousesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// offset
	Offset *int32 `json:"offset,omitempty"`

	// limit
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSecurityDlfDataWareHousesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDlfDataWareHousesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityDlfDataWareHousesRequest", string(data)}, " ")
}
