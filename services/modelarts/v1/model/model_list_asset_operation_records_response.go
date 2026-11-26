package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetOperationRecordsResponse Response Object
type ListAssetOperationRecordsResponse struct {

	// **参数解释**：操作记录总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：模型操作记录。
	Data           *[]AssetOperationRecordRsp `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListAssetOperationRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetOperationRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListAssetOperationRecordsResponse", string(data)}, " ")
}
