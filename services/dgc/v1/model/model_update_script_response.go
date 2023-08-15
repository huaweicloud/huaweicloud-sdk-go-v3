package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScriptResponse Response Object
type UpdateScriptResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateScriptResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScriptResponse struct{}"
	}

	return strings.Join([]string{"UpdateScriptResponse", string(data)}, " ")
}
