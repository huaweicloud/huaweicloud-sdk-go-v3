package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteServiceSetUsingDeleteResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteServiceSetUsingDeleteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceSetUsingDeleteResponse struct{}"
	}

	return strings.Join([]string{"DeleteServiceSetUsingDeleteResponse", string(data)}, " ")
}
