package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudConnectionResponse Response Object
type DeleteCloudConnectionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCloudConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudConnectionResponse struct{}"
	}

	return strings.Join([]string{"DeleteCloudConnectionResponse", string(data)}, " ")
}
