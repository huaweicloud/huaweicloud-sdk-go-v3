package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RunTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTaskResponse struct{}"
	}

	return strings.Join([]string{"RunTaskResponse", string(data)}, " ")
}
