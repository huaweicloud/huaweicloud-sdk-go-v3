package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDiskInfoResponse Response Object
type UpdateDiskInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDiskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDiskInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateDiskInfoResponse", string(data)}, " ")
}
