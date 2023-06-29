package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateArchiveConfigResponse Response Object
type UpdateArchiveConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateArchiveConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateArchiveConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateArchiveConfigResponse", string(data)}, " ")
}
