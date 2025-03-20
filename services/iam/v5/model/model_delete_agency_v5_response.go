package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgencyV5Response Response Object
type DeleteAgencyV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAgencyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgencyV5Response struct{}"
	}

	return strings.Join([]string{"DeleteAgencyV5Response", string(data)}, " ")
}
