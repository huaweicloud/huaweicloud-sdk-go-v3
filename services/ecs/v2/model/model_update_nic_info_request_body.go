package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNicInfoRequestBody struct {
	InterfaceAttachment *InterfaceAttachment `json:"interface_attachment"`
}

func (o UpdateNicInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNicInfoRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNicInfoRequestBody", string(data)}, " ")
}
