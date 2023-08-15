package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMetadataReqGesMetadata 存储metadata的消息信息的对象。
type CreateMetadataReqGesMetadata struct {

	// label列表
	Labels []CreateMetadataReqGesMetadataLabels `json:"labels"`
}

func (o CreateMetadataReqGesMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetadataReqGesMetadata struct{}"
	}

	return strings.Join([]string{"CreateMetadataReqGesMetadata", string(data)}, " ")
}
