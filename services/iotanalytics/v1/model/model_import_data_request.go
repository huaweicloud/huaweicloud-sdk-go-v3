package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportDataRequest Request Object
type ImportDataRequest struct {
	Body *ImportDataRequestBody `json:"body,omitempty"`
}

func (o ImportDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataRequest struct{}"
	}

	return strings.Join([]string{"ImportDataRequest", string(data)}, " ")
}
