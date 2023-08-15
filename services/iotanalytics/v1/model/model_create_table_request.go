package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableRequest Request Object
type CreateTableRequest struct {
	Body *CreateTableRequestBody `json:"body,omitempty"`
}

func (o CreateTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableRequest struct{}"
	}

	return strings.Join([]string{"CreateTableRequest", string(data)}, " ")
}
