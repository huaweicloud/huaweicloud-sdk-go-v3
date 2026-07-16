package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkMetadataAnnotations 网络资源metadata信息中的annotations字段信息。
type NetworkMetadataAnnotations struct {

	// **参数解释**：网络的描述信息。 **取值范围**：不能包含字符!<>=&\"'。
	OsModelartsDescription *string `json:"os.modelarts/description,omitempty"`
}

func (o NetworkMetadataAnnotations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkMetadataAnnotations struct{}"
	}

	return strings.Join([]string{"NetworkMetadataAnnotations", string(data)}, " ")
}
