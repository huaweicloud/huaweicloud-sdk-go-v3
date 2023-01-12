package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// indicator batch operation response
type DataResponse struct {

	// id list
	SuccessIds []string `json:"success_ids"`

	// id list
	ErrorIds *[]string `json:"error_ids,omitempty"`
}

func (o DataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataResponse struct{}"
	}

	return strings.Join([]string{"DataResponse", string(data)}, " ")
}
