package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyClusterResponse Response Object
type ModifyClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyClusterResponse struct{}"
	}

	return strings.Join([]string{"ModifyClusterResponse", string(data)}, " ")
}
