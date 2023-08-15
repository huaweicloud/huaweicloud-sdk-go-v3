package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataStoreRequest Request Object
type UpdateDataStoreRequest struct {

	// 存储 ID
	DataStoreId string `json:"data_store_id"`

	Body *UpdateDataStore `json:"body,omitempty"`
}

func (o UpdateDataStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataStoreRequest struct{}"
	}

	return strings.Join([]string{"UpdateDataStoreRequest", string(data)}, " ")
}
