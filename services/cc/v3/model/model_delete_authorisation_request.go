package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAuthorisationRequest Request Object
type DeleteAuthorisationRequest struct {

	// 授权ID。
	Id string `json:"id"`
}

func (o DeleteAuthorisationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAuthorisationRequest struct{}"
	}

	return strings.Join([]string{"DeleteAuthorisationRequest", string(data)}, " ")
}
