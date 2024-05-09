package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrugDatabaseRequest Request Object
type CreateDrugDatabaseRequest struct {
	Body *CreateDatabaseReq2 `json:"body,omitempty"`
}

func (o CreateDrugDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateDrugDatabaseRequest", string(data)}, " ")
}
