package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMigprojectResponse Response Object
type UpdateMigprojectResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMigprojectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMigprojectResponse struct{}"
	}

	return strings.Join([]string{"UpdateMigprojectResponse", string(data)}, " ")
}
