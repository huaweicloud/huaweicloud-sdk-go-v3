package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetRecordSetsStatusResponse Response Object
type BatchSetRecordSetsStatusResponse struct {
	Links *Link `json:"links,omitempty"`

	// 设置record set的列表响应。
	Recordsets *[]RecordsetData `json:"recordsets,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchSetRecordSetsStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetRecordSetsStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchSetRecordSetsStatusResponse", string(data)}, " ")
}
