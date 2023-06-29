package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEdgeSiteResponse Response Object
type DeleteEdgeSiteResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEdgeSiteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeSiteResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeSiteResponse", string(data)}, " ")
}
