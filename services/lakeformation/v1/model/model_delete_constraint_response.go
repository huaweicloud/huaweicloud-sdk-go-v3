package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteConstraintResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConstraintResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConstraintResponse struct{}"
	}

	return strings.Join([]string{"DeleteConstraintResponse", string(data)}, " ")
}
