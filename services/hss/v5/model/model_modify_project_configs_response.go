package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyProjectConfigsResponse Response Object
type ModifyProjectConfigsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ModifyProjectConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyProjectConfigsResponse struct{}"
	}

	return strings.Join([]string{"ModifyProjectConfigsResponse", string(data)}, " ")
}
