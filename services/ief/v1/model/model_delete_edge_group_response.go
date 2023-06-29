package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEdgeGroupResponse Response Object
type DeleteEdgeGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEdgeGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeGroupResponse", string(data)}, " ")
}
