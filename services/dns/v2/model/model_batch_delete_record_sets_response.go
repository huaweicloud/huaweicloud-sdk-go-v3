package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteRecordSetsResponse struct {

	// 待删除zone类型，当前仅支持public或private。
	Recordsets *[]RecordsetData `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteRecordSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRecordSetsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteRecordSetsResponse", string(data)}, " ")
}
