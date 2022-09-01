package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DebugCaseReturnHeader struct {

	// Connection
	Connection *string `json:"Connection,omitempty" xml:"Connection"`

	// Content-Length
	ContentLength *string `json:"Content-Length,omitempty" xml:"Content-Length"`

	// Content-Type
	ContentType *string `json:"Content-Type,omitempty" xml:"Content-Type"`

	// Date
	Date *string `json:"Date,omitempty" xml:"Date"`

	// Vary
	Vary *string `json:"Vary,omitempty" xml:"Vary"`
}

func (o DebugCaseReturnHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugCaseReturnHeader struct{}"
	}

	return strings.Join([]string{"DebugCaseReturnHeader", string(data)}, " ")
}
