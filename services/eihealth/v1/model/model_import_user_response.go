package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportUserResponse Response Object
type ImportUserResponse struct {
	Body           *[]UserIdRsp `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ImportUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportUserResponse struct{}"
	}

	return strings.Join([]string{"ImportUserResponse", string(data)}, " ")
}
