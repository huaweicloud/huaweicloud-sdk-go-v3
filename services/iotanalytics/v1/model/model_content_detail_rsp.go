package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTA服务各类数据源详细配置内容
type ContentDetailRsp struct {
	IotdaContent *IotdaContentRsp `json:"iotda_content,omitempty" xml:"iotda_content"`

	ObsContent *ObsContentRsp `json:"obs_content,omitempty" xml:"obs_content"`

	DisContent *DisContentRsp `json:"dis_content,omitempty" xml:"dis_content"`

	SmnContent *SmnContentRsp `json:"smn_content,omitempty" xml:"smn_content"`

	FunctionGraphContent *FunctionGraphContentRsp `json:"function_graph_content,omitempty" xml:"function_graph_content"`

	ModelArtsContent *ModelArtsContentRsp `json:"model_arts_content,omitempty" xml:"model_arts_content"`

	DcsContent *DcsContentRsp `json:"dcs_content,omitempty" xml:"dcs_content"`

	KafkaContent *KafkaContentRsp `json:"kafka_content,omitempty" xml:"kafka_content"`

	ApiContent *ApiContentRsp `json:"api_content,omitempty" xml:"api_content"`

	NodeContent *NodeContentRsp `json:"node_content,omitempty" xml:"node_content"`

	EdgeContent *EdgeContentRsp `json:"edge_content,omitempty" xml:"edge_content"`
}

func (o ContentDetailRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentDetailRsp struct{}"
	}

	return strings.Join([]string{"ContentDetailRsp", string(data)}, " ")
}
