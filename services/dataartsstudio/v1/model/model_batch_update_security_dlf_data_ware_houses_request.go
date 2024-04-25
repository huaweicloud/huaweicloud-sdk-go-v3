package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateSecurityDlfDataWareHousesRequest Request Object
type BatchUpdateSecurityDlfDataWareHousesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *FgacUpdateReq `json:"body,omitempty"`
}

func (o BatchUpdateSecurityDlfDataWareHousesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateSecurityDlfDataWareHousesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateSecurityDlfDataWareHousesRequest", string(data)}, " ")
}
