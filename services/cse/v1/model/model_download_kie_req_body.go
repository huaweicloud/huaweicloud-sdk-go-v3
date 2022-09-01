package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownloadKieReqBody struct {

	// 配置ID的集合
	Ids []string `json:"ids" xml:"ids"`
}

func (o DownloadKieReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKieReqBody struct{}"
	}

	return strings.Join([]string{"DownloadKieReqBody", string(data)}, " ")
}
