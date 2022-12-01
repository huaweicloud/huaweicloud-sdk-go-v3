package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckClusterResponse struct{}"
	}

	return strings.Join([]string{"CheckClusterResponse", string(data)}, " ")
}
