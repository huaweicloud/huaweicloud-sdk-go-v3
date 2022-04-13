package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTA服务各类数据源详细配置内容
type ContentDetailRsp struct {
	IotdaContent *IotdaContentRsp `json:"iotda_content,omitempty"`

	ObsContent *ObsContentRsp `json:"obs_content,omitempty"`

	DisContent *DisContentRsp `json:"dis_content,omitempty"`

	SmnContent *SmnContentRsp `json:"smn_content,omitempty"`

	FunctionGraphContent *FunctionGraphContentRsp `json:"function_graph_content,omitempty"`

	ModelArtsContent *ModelArtsContentRsp `json:"model_arts_content,omitempty"`

	DcsContent *DcsContentRsp `json:"dcs_content,omitempty"`

	KafkaContent *KafkaContentRsp `json:"kafka_content,omitempty"`

	ApiContent *ApiContentRsp `json:"api_content,omitempty"`

	NodeContent *NodeContentRsp `json:"node_content,omitempty"`

	EdgeContent *EdgeContentRsp `json:"edge_content,omitempty"`
}

func (o ContentDetailRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentDetailRsp struct{}"
	}

	return strings.Join([]string{"ContentDetailRsp", string(data)}, " ")
}
