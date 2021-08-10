package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRouteTableRequest struct {
	// 路由表ID

	RoutetableId string `json:"routetable_id"`
}

func (o DeleteRouteTableRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRouteTableRequest struct{}"
	}

	return strings.Join([]string{"DeleteRouteTableRequest", string(data)}, " ")
}
