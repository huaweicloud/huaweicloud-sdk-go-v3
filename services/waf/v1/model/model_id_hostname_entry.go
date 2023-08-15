package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdHostnameEntry 独享引擎实例防护域名信息
type IdHostnameEntry struct {

	// 防护域名ID
	Id string `json:"id"`

	// 防护域名
	Hostname string `json:"hostname"`
}

func (o IdHostnameEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdHostnameEntry struct{}"
	}

	return strings.Join([]string{"IdHostnameEntry", string(data)}, " ")
}
