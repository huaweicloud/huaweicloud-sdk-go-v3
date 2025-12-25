package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetrievalResultInfo 检索结果。
type RetrievalResultInfo struct {

	// **参数解释**： 文件ID（或FAQ ID）。  **取值范围**： 不涉及。
	FileId *string `json:"file_id,omitempty"`

	// **参数解释**： 文档标题（如果是FAQ，返回QUESTION）。  **取值范围**： 不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释**： 分片ID。  **取值范围**： 不涉及。
	ChunkId *string `json:"chunk_id,omitempty"`

	// **参数解释**： 文本内容（如果是FAQ，返回ANSWER）。  **取值范围**： 不涉及。
	Content *string `json:"content,omitempty"`

	// **参数解释**： 相似度。  **取值范围**： [0.0, 1.0]，包含两端。
	Similarity *float32 `json:"similarity,omitempty"`

	// **参数解释**： 知识库ID。  **取值范围**： 不涉及。
	KnowledgeBaseId *string `json:"knowledge_base_id,omitempty"`

	// **参数解释**： 检索到的图片列表（需知识库支持），与content中的图片占位符{KI|N}保持一一对应关系，N为图片索引值，从0开始。  **取值范围**： 不涉及。
	ImageIds *[]string `json:"image_ids,omitempty"`
}

func (o RetrievalResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetrievalResultInfo struct{}"
	}

	return strings.Join([]string{"RetrievalResultInfo", string(data)}, " ")
}
