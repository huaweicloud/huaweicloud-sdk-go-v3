package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateConnectionRequestBody struct {
	Connection *UpdateConnectionOption `json:"connection"`
}

func (o UpdateConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateConnectionRequestBody", string(data)}, " ")
}
