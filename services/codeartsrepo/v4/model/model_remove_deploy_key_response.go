package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveDeployKeyResponse Response Object
type RemoveDeployKeyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveDeployKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveDeployKeyResponse struct{}"
	}

	return strings.Join([]string{"RemoveDeployKeyResponse", string(data)}, " ")
}
