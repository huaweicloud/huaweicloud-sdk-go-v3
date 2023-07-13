package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateParaGroupNameResponse Response Object
type ValidateParaGroupNameResponse struct {

	// 校验结果。true为已存在，false为不存在。
	Exist          *bool `json:"exist,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ValidateParaGroupNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateParaGroupNameResponse struct{}"
	}

	return strings.Join([]string{"ValidateParaGroupNameResponse", string(data)}, " ")
}
