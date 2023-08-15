package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetadataRespGesMetadata 存储metadata的消息信息的对象。
type ShowMetadataRespGesMetadata struct {

	// Label数据结构集合。
	Labels *[]ShowMetadataRespGesMetadataLabels `json:"labels,omitempty"`
}

func (o ShowMetadataRespGesMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetadataRespGesMetadata struct{}"
	}

	return strings.Join([]string{"ShowMetadataRespGesMetadata", string(data)}, " ")
}
