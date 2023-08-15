package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddProtectBranchV2Response Response Object
type AddProtectBranchV2Response struct {
	Error *Error `json:"error,omitempty"`

	Result *AddProtectResponse `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddProtectBranchV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectBranchV2Response struct{}"
	}

	return strings.Join([]string{"AddProtectBranchV2Response", string(data)}, " ")
}
