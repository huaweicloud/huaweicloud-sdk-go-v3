package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeletePtrRecordsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeletePtrRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePtrRecordsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePtrRecordsResponse", string(data)}, " ")
}
