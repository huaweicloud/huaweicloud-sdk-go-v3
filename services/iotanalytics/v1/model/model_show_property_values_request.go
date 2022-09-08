package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPropertyValuesRequest struct {

	// 存储ID
	DataStoreId string `json:"data_store_id"`

	Body *GetPropertyRequest `json:"body,omitempty"`
}

func (o ShowPropertyValuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPropertyValuesRequest struct{}"
	}

	return strings.Join([]string{"ShowPropertyValuesRequest", string(data)}, " ")
}
