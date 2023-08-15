package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMessageEmailConfigResponse Response Object
type DeleteMessageEmailConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMessageEmailConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMessageEmailConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteMessageEmailConfigResponse", string(data)}, " ")
}
