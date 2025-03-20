package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgencyV5Response Response Object
type UpdateAgencyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAgencyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyV5Response struct{}"
	}

	return strings.Join([]string{"UpdateAgencyV5Response", string(data)}, " ")
}
