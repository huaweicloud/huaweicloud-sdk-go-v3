package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreHbaInfoResponse Response Object
type RestoreHbaInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreHbaInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreHbaInfoResponse struct{}"
	}

	return strings.Join([]string{"RestoreHbaInfoResponse", string(data)}, " ")
}
