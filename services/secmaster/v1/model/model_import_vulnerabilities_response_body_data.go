package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportVulnerabilitiesResponseBodyData 漏洞导入结果返回数据
type ImportVulnerabilitiesResponseBodyData struct {

	// 导入的漏洞id列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o ImportVulnerabilitiesResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportVulnerabilitiesResponseBodyData struct{}"
	}

	return strings.Join([]string{"ImportVulnerabilitiesResponseBodyData", string(data)}, " ")
}
