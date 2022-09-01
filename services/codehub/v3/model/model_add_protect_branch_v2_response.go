package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddProtectBranchV2Response struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *AddProtectResponse `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o AddProtectBranchV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectBranchV2Response struct{}"
	}

	return strings.Join([]string{"AddProtectBranchV2Response", string(data)}, " ")
}
