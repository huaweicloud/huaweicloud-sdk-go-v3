package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteLabelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteLabelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLabelResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteLabelResponse", string(data)}, " ")
}
