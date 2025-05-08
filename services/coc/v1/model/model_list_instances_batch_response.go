package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesBatchResponse Response Object
type ListInstancesBatchResponse struct {

	// 分批结果
	Data           *[]InstancesBatchResultMode `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListInstancesBatchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesBatchResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesBatchResponse", string(data)}, " ")
}
