package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNodeRequestBody struct {
	Node *UpdateNodeReqDetail `json:"node"`
}

func (o UpdateNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNodeRequestBody", string(data)}, " ")
}
