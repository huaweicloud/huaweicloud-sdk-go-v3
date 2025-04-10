package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportVulRequestBody struct {

	// 导出漏洞数据的表头信息列表
	ExportHeaders *[][]string `json:"export_headers,omitempty"`

	// 指定导出的漏洞id列表
	VulIdList *[]string `json:"vul_id_list,omitempty"`

	// 指定导出的主机id列表
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o ExportVulRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportVulRequestBody struct{}"
	}

	return strings.Join([]string{"ExportVulRequestBody", string(data)}, " ")
}
