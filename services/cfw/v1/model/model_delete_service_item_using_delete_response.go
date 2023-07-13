package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServiceItemUsingDeleteResponse Response Object
type DeleteServiceItemUsingDeleteResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteServiceItemUsingDeleteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServiceItemUsingDeleteResponse struct{}"
	}

	return strings.Join([]string{"DeleteServiceItemUsingDeleteResponse", string(data)}, " ")
}
