package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIsolatedFileRequestInfo 删除已隔离的文件详情
type DeleteIsolatedFileRequestInfo struct {

	// 需要删除的文件列表
	DataList []DelIsolatedFileRequestInfo `json:"data_list"`
}

func (o DeleteIsolatedFileRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIsolatedFileRequestInfo struct{}"
	}

	return strings.Join([]string{"DeleteIsolatedFileRequestInfo", string(data)}, " ")
}
