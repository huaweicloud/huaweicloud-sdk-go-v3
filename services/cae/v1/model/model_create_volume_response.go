package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVolumeResponse Response Object
type CreateVolumeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeResponse struct{}"
	}

	return strings.Join([]string{"CreateVolumeResponse", string(data)}, " ")
}
