package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddScanJobRequest Request Object
type AddScanJobRequest struct {
	Body *ScanJobRequest `json:"body,omitempty"`
}

func (o AddScanJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddScanJobRequest struct{}"
	}

	return strings.Join([]string{"AddScanJobRequest", string(data)}, " ")
}
