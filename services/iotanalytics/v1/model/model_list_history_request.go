package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHistoryRequest struct {
	// 存储ID

	DataStoreId string `json:"data_store_id"`

	Body *GetHistoryRequest `json:"body,omitempty"`
}

func (o ListHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryRequest", string(data)}, " ")
}
