package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePersistentStorageRequest Request Object
type CreatePersistentStorageRequest struct {
	Body *CreatePersistentStorageReq `json:"body,omitempty"`
}

func (o CreatePersistentStorageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePersistentStorageRequest struct{}"
	}

	return strings.Join([]string{"CreatePersistentStorageRequest", string(data)}, " ")
}
