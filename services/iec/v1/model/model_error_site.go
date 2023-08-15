package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ErrorSite 异常站点
type ErrorSite struct {

	// 异常站点。
	ErrSites *[]string `json:"err_sites,omitempty"`
}

func (o ErrorSite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorSite struct{}"
	}

	return strings.Join([]string{"ErrorSite", string(data)}, " ")
}
