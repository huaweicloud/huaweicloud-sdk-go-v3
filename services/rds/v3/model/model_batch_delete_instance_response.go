package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteInstanceResponse Response Object
type BatchDeleteInstanceResponse struct {

	// 删除实例记录
	Records        *[]CustomerBatchDeleteInstanceRecord `json:"records,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o BatchDeleteInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceResponse", string(data)}, " ")
}
