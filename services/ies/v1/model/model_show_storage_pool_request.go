package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowStoragePoolRequest struct {

	// 存储池ID
	Id string `json:"id"`
}

func (o ShowStoragePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStoragePoolRequest struct{}"
	}

	return strings.Join([]string{"ShowStoragePoolRequest", string(data)}, " ")
}
