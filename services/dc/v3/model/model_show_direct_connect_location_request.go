package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDirectConnectLocationRequest Request Object
type ShowDirectConnectLocationRequest struct {

	// direct-connect-locationID。
	DirectConnectLocationId string `json:"direct_connect_location_id"`
}

func (o ShowDirectConnectLocationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDirectConnectLocationRequest struct{}"
	}

	return strings.Join([]string{"ShowDirectConnectLocationRequest", string(data)}, " ")
}
