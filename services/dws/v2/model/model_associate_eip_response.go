package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AssociateEipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateEipResponse struct{}"
	}

	return strings.Join([]string{"AssociateEipResponse", string(data)}, " ")
}
