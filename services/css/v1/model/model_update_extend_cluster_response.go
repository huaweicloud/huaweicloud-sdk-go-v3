package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExtendClusterResponse Response Object
type UpdateExtendClusterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateExtendClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExtendClusterResponse struct{}"
	}

	return strings.Join([]string{"UpdateExtendClusterResponse", string(data)}, " ")
}
