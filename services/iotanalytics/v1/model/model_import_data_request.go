package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportDataRequest struct {
	Body *ImportDataRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ImportDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataRequest struct{}"
	}

	return strings.Join([]string{"ImportDataRequest", string(data)}, " ")
}
