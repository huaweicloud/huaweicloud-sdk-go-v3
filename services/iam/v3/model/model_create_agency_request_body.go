package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgencyRequestBody
type CreateAgencyRequestBody struct {
	Agency *CreateAgencyOption `json:"agency"`
}

func (o CreateAgencyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAgencyRequestBody", string(data)}, " ")
}
