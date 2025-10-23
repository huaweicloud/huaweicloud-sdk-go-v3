package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveDeployKeyFromSubmodulesResponse Response Object
type RemoveDeployKeyFromSubmodulesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveDeployKeyFromSubmodulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveDeployKeyFromSubmodulesResponse struct{}"
	}

	return strings.Join([]string{"RemoveDeployKeyFromSubmodulesResponse", string(data)}, " ")
}
