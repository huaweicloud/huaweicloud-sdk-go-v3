package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ExportKnowledgeSkillResponse Response Object
type ExportKnowledgeSkillResponse struct {
	XRequestId     *string       `json:"X-Request-Id,omitempty"`
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportKnowledgeSkillResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportKnowledgeSkillResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportKnowledgeSkillResponse struct{}"
	}

	return strings.Join([]string{"ExportKnowledgeSkillResponse", string(data)}, " ")
}
