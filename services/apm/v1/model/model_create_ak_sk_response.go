package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAkSkResponse Response Object
type CreateAkSkResponse struct {

	// 创建/删除的ak信息。
	Ak *string `json:"ak,omitempty"`

	// 创建/删除的sk信息。
	Sk             *string `json:"sk,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAkSkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAkSkResponse struct{}"
	}

	return strings.Join([]string{"CreateAkSkResponse", string(data)}, " ")
}
