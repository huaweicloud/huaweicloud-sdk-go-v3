package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublicScriptsResponse Response Object
type ListPublicScriptsResponse struct {
	Data           *PublicScriptListPage `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListPublicScriptsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicScriptsResponse struct{}"
	}

	return strings.Join([]string{"ListPublicScriptsResponse", string(data)}, " ")
}
