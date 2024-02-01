package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountGcbResourceByTagRequest Request Object
type CountGcbResourceByTagRequest struct {
	Body *QueryResourceByTagRequestBody `json:"body,omitempty"`
}

func (o CountGcbResourceByTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountGcbResourceByTagRequest struct{}"
	}

	return strings.Join([]string{"CountGcbResourceByTagRequest", string(data)}, " ")
}
