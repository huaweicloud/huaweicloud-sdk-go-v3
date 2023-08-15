package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlInstanceOpsWindowResponse Response Object
type UpdateGaussMySqlInstanceOpsWindowResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlInstanceOpsWindowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceOpsWindowResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceOpsWindowResponse", string(data)}, " ")
}
