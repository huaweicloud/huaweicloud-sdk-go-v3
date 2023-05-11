package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateFsDirResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateFsDirResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFsDirResponse struct{}"
	}

	return strings.Join([]string{"CreateFsDirResponse", string(data)}, " ")
}
