package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePoolCascadeResponse Response Object
type DeletePoolCascadeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePoolCascadeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePoolCascadeResponse struct{}"
	}

	return strings.Join([]string{"DeletePoolCascadeResponse", string(data)}, " ")
}
