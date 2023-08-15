package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStoragePoolRequest Request Object
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
