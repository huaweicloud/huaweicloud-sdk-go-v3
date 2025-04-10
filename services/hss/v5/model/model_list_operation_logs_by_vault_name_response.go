package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOperationLogsByVaultNameResponse Response Object
type ListOperationLogsByVaultNameResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// list
	DataList       *[]OperationLogInfo `json:"data_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListOperationLogsByVaultNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOperationLogsByVaultNameResponse struct{}"
	}

	return strings.Join([]string{"ListOperationLogsByVaultNameResponse", string(data)}, " ")
}
