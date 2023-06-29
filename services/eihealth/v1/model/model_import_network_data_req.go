package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportNetworkDataReq 导入网上数据请求体
type ImportNetworkDataReq struct {

	// 所在文件夹
	TargetFolder *string `json:"target_folder,omitempty"`

	// 导入网上数据的url集
	Urls []string `json:"urls"`

	// 导入网上数据的md5集
	Md5s *[]string `json:"md5s,omitempty"`
}

func (o ImportNetworkDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportNetworkDataReq struct{}"
	}

	return strings.Join([]string{"ImportNetworkDataReq", string(data)}, " ")
}
