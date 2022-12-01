package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEdgeGroupCertResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEdgeGroupCertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeGroupCertResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeGroupCertResponse", string(data)}, " ")
}
