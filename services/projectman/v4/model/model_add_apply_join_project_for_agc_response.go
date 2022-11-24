package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddApplyJoinProjectForAgcResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddApplyJoinProjectForAgcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddApplyJoinProjectForAgcResponse struct{}"
	}

	return strings.Join([]string{"AddApplyJoinProjectForAgcResponse", string(data)}, " ")
}
