package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDesignTableModelsFromLogicResponse Response Object
type BatchCreateDesignTableModelsFromLogicResponse struct {
	Data           *CreateWorkspaceResultData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o BatchCreateDesignTableModelsFromLogicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDesignTableModelsFromLogicResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateDesignTableModelsFromLogicResponse", string(data)}, " ")
}
