package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourcesLevelResponse Response Object
type CreateResourcesLevelResponse struct {

	// 资源分级ID
	ResourceLevelId *string `json:"resource_level_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateResourcesLevelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourcesLevelResponse struct{}"
	}

	return strings.Join([]string{"CreateResourcesLevelResponse", string(data)}, " ")
}
