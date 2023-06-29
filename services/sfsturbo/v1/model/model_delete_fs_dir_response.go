package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFsDirResponse Response Object
type DeleteFsDirResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteFsDirResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFsDirResponse struct{}"
	}

	return strings.Join([]string{"DeleteFsDirResponse", string(data)}, " ")
}
