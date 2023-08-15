package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CinderDeleteVolumeTransferResponse Response Object
type CinderDeleteVolumeTransferResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CinderDeleteVolumeTransferResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderDeleteVolumeTransferResponse struct{}"
	}

	return strings.Join([]string{"CinderDeleteVolumeTransferResponse", string(data)}, " ")
}
