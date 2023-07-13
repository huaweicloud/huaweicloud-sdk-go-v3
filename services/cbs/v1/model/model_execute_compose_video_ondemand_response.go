package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteComposeVideoOndemandResponse Response Object
type ExecuteComposeVideoOndemandResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExecuteComposeVideoOndemandResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteComposeVideoOndemandResponse struct{}"
	}

	return strings.Join([]string{"ExecuteComposeVideoOndemandResponse", string(data)}, " ")
}
