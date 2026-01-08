package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportIpRequest Request Object
type ImportIpRequest struct {
	Body *ImportIpRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportIpRequest struct{}"
	}

	return strings.Join([]string{"ImportIpRequest", string(data)}, " ")
}
