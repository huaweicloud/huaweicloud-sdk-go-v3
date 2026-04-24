package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebuildDdmConfigResponse Response Object
type RebuildDdmConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RebuildDdmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildDdmConfigResponse struct{}"
	}

	return strings.Join([]string{"RebuildDdmConfigResponse", string(data)}, " ")
}
