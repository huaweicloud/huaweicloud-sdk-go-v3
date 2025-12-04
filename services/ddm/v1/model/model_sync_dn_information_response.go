package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncDnInformationResponse Response Object
type SyncDnInformationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SyncDnInformationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncDnInformationResponse struct{}"
	}

	return strings.Join([]string{"SyncDnInformationResponse", string(data)}, " ")
}
