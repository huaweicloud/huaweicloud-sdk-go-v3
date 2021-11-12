package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListScriptsResponse struct {
	Total *int32 `json:"total,omitempty"`

	Scripts        *[]ScriptInfo `json:"scripts,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListScriptsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptsResponse struct{}"
	}

	return strings.Join([]string{"ListScriptsResponse", string(data)}, " ")
}
