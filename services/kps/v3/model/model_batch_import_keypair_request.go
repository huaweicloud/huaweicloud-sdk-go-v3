package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchImportKeypairRequest Request Object
type BatchImportKeypairRequest struct {
	Body *[]CreateKeypairRequestBody `json:"body,omitempty"`
}

func (o BatchImportKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportKeypairRequest struct{}"
	}

	return strings.Join([]string{"BatchImportKeypairRequest", string(data)}, " ")
}
