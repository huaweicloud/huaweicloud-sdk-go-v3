package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteBaremetalServerTagsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteBaremetalServerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBaremetalServerTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteBaremetalServerTagsResponse", string(data)}, " ")
}
