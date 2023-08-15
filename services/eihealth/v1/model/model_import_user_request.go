package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportUserRequest Request Object
type ImportUserRequest struct {
	Body *ImportUserReq `json:"body,omitempty"`
}

func (o ImportUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportUserRequest struct{}"
	}

	return strings.Join([]string{"ImportUserRequest", string(data)}, " ")
}
