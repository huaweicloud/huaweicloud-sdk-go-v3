package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConnectionRequestBody struct {
	Connection *CreateConnectionOption `json:"connection"`
}

func (o CreateConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConnectionRequestBody", string(data)}, " ")
}
