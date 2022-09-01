package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DebugCaseResult struct {

	// body
	Body *string `json:"body,omitempty" xml:"body"`

	// errorReason
	ErrorReason *string `json:"errorReason,omitempty" xml:"errorReason"`

	Header *DebugCaseResultHeader `json:"header,omitempty" xml:"header"`

	// name
	Name *string `json:"name,omitempty" xml:"name"`

	// responseTime
	ResponseTime *int32 `json:"responseTime,omitempty" xml:"responseTime"`

	// result
	Result *int32 `json:"result,omitempty" xml:"result"`

	// returnBody
	ReturnBody *string `json:"returnBody,omitempty" xml:"returnBody"`

	ReturnHeader *DebugCaseReturnHeader `json:"returnHeader,omitempty" xml:"returnHeader"`

	// statusCode
	StatusCode *string `json:"statusCode,omitempty" xml:"statusCode"`

	// url
	Url *string `json:"url,omitempty" xml:"url"`
}

func (o DebugCaseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugCaseResult struct{}"
	}

	return strings.Join([]string{"DebugCaseResult", string(data)}, " ")
}
