package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAppQuotaBindingApp struct {

	// 客户端应用编号列表
	AppIds []string `json:"app_ids" xml:"app_ids"`
}

func (o CreateAppQuotaBindingApp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppQuotaBindingApp struct{}"
	}

	return strings.Join([]string{"CreateAppQuotaBindingApp", string(data)}, " ")
}
