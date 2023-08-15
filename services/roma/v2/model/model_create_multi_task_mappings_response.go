package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMultiTaskMappingsResponse Response Object
type CreateMultiTaskMappingsResponse struct {

	// 映射唯一ID
	MappingId      *string `json:"mapping_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMultiTaskMappingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMultiTaskMappingsResponse struct{}"
	}

	return strings.Join([]string{"CreateMultiTaskMappingsResponse", string(data)}, " ")
}
