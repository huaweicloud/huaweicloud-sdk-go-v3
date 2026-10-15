package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobStateResponse Response Object
type ShowSparkJobStateResponse struct {

	// **参数解释**：Spark作业ID，用于唯一标识一个Spark作业。 **取值范围**：采用UUID格式，长度为36个字符，例如：6c98db52-cac2-4ff1-9a91-b7793e95557d。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：Spark作业的状态。 **取值范围**： - QUEUED：排队中。 - PENDING：启动中。 - RUNNING：运行中。 - CANCELING：取消中。 - CANCELED：已取消。 - FAILED：运行失败。 - QUEUED_TIMEOUT：排队超时。 - RUNNING_TIMEOUT：运行超时。 - SUCCEED：运行成功。
	State          *string `json:"state,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSparkJobStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobStateResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJobStateResponse", string(data)}, " ")
}
