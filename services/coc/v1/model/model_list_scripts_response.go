package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptsResponse Response Object
type ListScriptsResponse struct {
	Data           *ScriptListPage `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListScriptsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptsResponse struct{}"
	}

	return strings.Join([]string{"ListScriptsResponse", string(data)}, " ")
}
