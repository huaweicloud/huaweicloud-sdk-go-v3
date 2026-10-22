package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobUpdateRequestBodyMetadata **参数解释**：训练作业元信息。 **约束限制**：不涉及。
type JobUpdateRequestBodyMetadata struct {

	// **参数解释**：训练作业高级功能配置。可选取值如下：   - \"jupyter-lab/enable\": \"true\"（JupyterLab训练应用程序） **约束限制**：不涉及。
	Annotation map[string]string `json:"annotation,omitempty"`
}

func (o JobUpdateRequestBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobUpdateRequestBodyMetadata struct{}"
	}

	return strings.Join([]string{"JobUpdateRequestBodyMetadata", string(data)}, " ")
}
