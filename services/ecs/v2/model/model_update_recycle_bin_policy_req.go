package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRecycleBinPolicyReq struct {
	Policy *UpdateRecycleBinPolicyOption `json:"policy"`
}

func (o UpdateRecycleBinPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinPolicyReq struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinPolicyReq", string(data)}, " ")
}
