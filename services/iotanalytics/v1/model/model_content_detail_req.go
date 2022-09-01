package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTA服务各类数据源详细配置内容
type ContentDetailReq struct {
	IotdaContent *IotdaContentReq `json:"iotda_content,omitempty" xml:"iotda_content"`

	ObsContent *ObsContentReq `json:"obs_content,omitempty" xml:"obs_content"`

	DisContent *DisContentReq `json:"dis_content,omitempty" xml:"dis_content"`

	SmnContent *SmnContentReq `json:"smn_content,omitempty" xml:"smn_content"`

	FunctionGraphContent *FunctionGraphContentReq `json:"function_graph_content,omitempty" xml:"function_graph_content"`

	ModelArtsContent *ModelArtsContentReq `json:"model_arts_content,omitempty" xml:"model_arts_content"`

	DcsContent *DcsContentReq `json:"dcs_content,omitempty" xml:"dcs_content"`

	KafkaContent *KafkaContentReq `json:"kafka_content,omitempty" xml:"kafka_content"`

	ApiContent *ApiContentReq `json:"api_content,omitempty" xml:"api_content"`

	NodeContent *NodeContentReq `json:"node_content,omitempty" xml:"node_content"`

	EdgeContent *EdgeContentReq `json:"edge_content,omitempty" xml:"edge_content"`
}

func (o ContentDetailReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentDetailReq struct{}"
	}

	return strings.Join([]string{"ContentDetailReq", string(data)}, " ")
}
