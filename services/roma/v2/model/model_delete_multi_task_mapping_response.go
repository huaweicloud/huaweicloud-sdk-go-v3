package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMultiTaskMappingResponse Response Object
type DeleteMultiTaskMappingResponse struct {

	// 映射唯一ID
	MappingId      *string `json:"mapping_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMultiTaskMappingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMultiTaskMappingResponse struct{}"
	}

	return strings.Join([]string{"DeleteMultiTaskMappingResponse", string(data)}, " ")
}
