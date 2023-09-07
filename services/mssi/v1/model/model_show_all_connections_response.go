package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAllConnectionsResponse Response Object
type ShowAllConnectionsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ShowAllConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ShowAllConnectionsResponse", string(data)}, " ")
}
