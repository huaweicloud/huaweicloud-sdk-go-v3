package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyTemplateResponse Response Object
type DeletePolicyTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePolicyTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeletePolicyTemplateResponse", string(data)}, " ")
}
