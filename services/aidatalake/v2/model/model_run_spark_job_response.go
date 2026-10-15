package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSparkJobResponse Response Object
type RunSparkJobResponse struct {

	// **参数解释**：Spark作业ID，用于唯一标识本次启动的Spark作业。 **取值范围**：只能由十六进制字符、中划线及下划线组成，长度为1~64个字符，例如：1ee7d9db-90ec-4b71-89b8-0de936dade77。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunSparkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSparkJobResponse struct{}"
	}

	return strings.Join([]string{"RunSparkJobResponse", string(data)}, " ")
}
