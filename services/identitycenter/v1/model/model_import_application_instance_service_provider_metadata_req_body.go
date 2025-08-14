package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportApplicationInstanceServiceProviderMetadataReqBody ImportApplicationInstanceServiceProviderMetadata的请求体
type ImportApplicationInstanceServiceProviderMetadataReqBody struct {

	// 元数据文件
	Metadata string `json:"metadata"`
}

func (o ImportApplicationInstanceServiceProviderMetadataReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportApplicationInstanceServiceProviderMetadataReqBody struct{}"
	}

	return strings.Join([]string{"ImportApplicationInstanceServiceProviderMetadataReqBody", string(data)}, " ")
}
