package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportTableRequest Request Object
type ImportTableRequest struct {
	Body *ImportTableRequestBody `json:"body,omitempty"`
}

func (o ImportTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportTableRequest struct{}"
	}

	return strings.Join([]string{"ImportTableRequest", string(data)}, " ")
}
