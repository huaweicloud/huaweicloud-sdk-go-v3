package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyRepositoryRequest Request Object
type ModifyRepositoryRequest struct {

	// tab_id
	TabId string `json:"tab_id"`

	Body *IdeRepositoryPair `json:"body,omitempty"`
}

func (o ModifyRepositoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyRepositoryRequest struct{}"
	}

	return strings.Join([]string{"ModifyRepositoryRequest", string(data)}, " ")
}
