package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchUpdateLakeFormationInstanceTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateLakeFormationInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateLakeFormationInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateLakeFormationInstanceTagsResponse", string(data)}, " ")
}
