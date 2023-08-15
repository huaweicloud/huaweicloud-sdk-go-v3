package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLakeFormationInstanceRequest Request Object
type CreateLakeFormationInstanceRequest struct {
	Body *CreateInstanceRequestBody `json:"body,omitempty"`
}

func (o CreateLakeFormationInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLakeFormationInstanceRequest struct{}"
	}

	return strings.Join([]string{"CreateLakeFormationInstanceRequest", string(data)}, " ")
}
