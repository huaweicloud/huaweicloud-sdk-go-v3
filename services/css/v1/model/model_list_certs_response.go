package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertsResponse Response Object
type ListCertsResponse struct {
	CertsRecords *CertsRecordsDatastore `json:"certsRecords,omitempty"`

	// 证书记录数量。
	TotalSize      *int32 `json:"totalSize,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCertsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertsResponse struct{}"
	}

	return strings.Join([]string{"ListCertsResponse", string(data)}, " ")
}
