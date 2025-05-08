package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDocumentAtomicsRequest Request Object
type ListDocumentAtomicsRequest struct {
	Limit *string `json:"limit,omitempty"`

	Offset *string `json:"offset,omitempty"`
}

func (o ListDocumentAtomicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDocumentAtomicsRequest struct{}"
	}

	return strings.Join([]string{"ListDocumentAtomicsRequest", string(data)}, " ")
}
