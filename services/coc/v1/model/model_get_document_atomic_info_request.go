package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDocumentAtomicInfoRequest Request Object
type GetDocumentAtomicInfoRequest struct {
	AtomicUniqueKey string `json:"atomic_unique_key"`
}

func (o GetDocumentAtomicInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDocumentAtomicInfoRequest struct{}"
	}

	return strings.Join([]string{"GetDocumentAtomicInfoRequest", string(data)}, " ")
}
