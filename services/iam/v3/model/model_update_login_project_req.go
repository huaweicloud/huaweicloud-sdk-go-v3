package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLoginProjectReq
type UpdateLoginProjectReq struct {
	LoginProtect *UpdateLoginProject `json:"login_protect"`
}

func (o UpdateLoginProjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoginProjectReq struct{}"
	}

	return strings.Join([]string{"UpdateLoginProjectReq", string(data)}, " ")
}
