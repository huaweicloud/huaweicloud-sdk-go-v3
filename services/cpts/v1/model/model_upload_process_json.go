package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadProcessJson json
type UploadProcessJson struct {

	// details
	Details *[]UploadProcessJsonDetail `json:"details,omitempty"`

	// process_status
	ProcessStatus *int32 `json:"process_status,omitempty"`
}

func (o UploadProcessJson) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadProcessJson struct{}"
	}

	return strings.Join([]string{"UploadProcessJson", string(data)}, " ")
}
