package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScopedToken
type ScopedToken struct {

	// 联邦unscoped token的ID。
	Id string `json:"id"`
}

func (o ScopedToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScopedToken struct{}"
	}

	return strings.Join([]string{"ScopedToken", string(data)}, " ")
}
