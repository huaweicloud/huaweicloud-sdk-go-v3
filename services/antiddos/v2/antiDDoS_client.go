package v2

import (
	httpclient "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/antiddos/v2/model"
)

type AntiDDoSClient struct {
	HcClient *httpclient.HcHttpClient
}

func NewAntiDDoSClient(hcClient *httpclient.HcHttpClient) *AntiDDoSClient {
	return &AntiDDoSClient{HcClient: hcClient}
}

func AntiDDoSClientBuilder() *httpclient.HcHttpClientBuilder {
	builder := httpclient.NewHcHttpClientBuilder()
	return builder
}

// ShowAlertConfig 查询告警配置信息
//
// 查询用户配置信息，用户可以通过此接口查询是否接收某类告警，同时可以配置是手机短信还是电子邮件接收告警信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AntiDDoSClient) ShowAlertConfig(request *model.ShowAlertConfigRequest) (*model.ShowAlertConfigResponse, error) {
	requestDef := GenReqDefForShowAlertConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertConfigResponse), nil
	}
}

// ShowAlertConfigInvoker 查询告警配置信息
func (c *AntiDDoSClient) ShowAlertConfigInvoker(request *model.ShowAlertConfigRequest) *ShowAlertConfigInvoker {
	requestDef := GenReqDefForShowAlertConfig()
	return &ShowAlertConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// UpdateAlertConfig 更新告警配置信息
//
// 更新用户配置信息，用户可以通过此接口更新是否接收某类告警，同时可以配置是手机短信还是电子邮件接收告警信息。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AntiDDoSClient) UpdateAlertConfig(request *model.UpdateAlertConfigRequest) (*model.UpdateAlertConfigResponse, error) {
	requestDef := GenReqDefForUpdateAlertConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlertConfigResponse), nil
	}
}

// UpdateAlertConfigInvoker 更新告警配置信息
func (c *AntiDDoSClient) UpdateAlertConfigInvoker(request *model.UpdateAlertConfigRequest) *UpdateAlertConfigInvoker {
	requestDef := GenReqDefForUpdateAlertConfig()
	return &UpdateAlertConfigInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListDDosStatus 查询EIP防护状态列表
//
// 查询用户所有EIP的Anti-DDoS防护状态信息，用户的EIP无论是否绑定到云服务器，都可以进行查询。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AntiDDoSClient) ListDDosStatus(request *model.ListDDosStatusRequest) (*model.ListDDosStatusResponse, error) {
	requestDef := GenReqDefForListDDosStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDDosStatusResponse), nil
	}
}

// ListDDosStatusInvoker 查询EIP防护状态列表
func (c *AntiDDoSClient) ListDDosStatusInvoker(request *model.ListDDosStatusRequest) *ListDDosStatusInvoker {
	requestDef := GenReqDefForListDDosStatus()
	return &ListDDosStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ListNewConfigs 查询Anti-DDoS配置可选范围
//
// 查询系统支持的Anti-DDoS防护策略配置的可选范围，用户根据范围列表选择适合自已业务的防护策略进行Anti-DDoS流量清洗。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AntiDDoSClient) ListNewConfigs(request *model.ListNewConfigsRequest) (*model.ListNewConfigsResponse, error) {
	requestDef := GenReqDefForListNewConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNewConfigsResponse), nil
	}
}

// ListNewConfigsInvoker 查询Anti-DDoS配置可选范围
func (c *AntiDDoSClient) ListNewConfigsInvoker(request *model.ListNewConfigsRequest) *ListNewConfigsInvoker {
	requestDef := GenReqDefForListNewConfigs()
	return &ListNewConfigsInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}

// ShowNewTaskStatus 查询Anti-DDoS任务
//
// 用户查询指定的Anti-DDoS防护配置任务，得到任务当前执行的状态。
//
// Please refer to HUAWEI cloud API Explorer for details.
func (c *AntiDDoSClient) ShowNewTaskStatus(request *model.ShowNewTaskStatusRequest) (*model.ShowNewTaskStatusResponse, error) {
	requestDef := GenReqDefForShowNewTaskStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNewTaskStatusResponse), nil
	}
}

// ShowNewTaskStatusInvoker 查询Anti-DDoS任务
func (c *AntiDDoSClient) ShowNewTaskStatusInvoker(request *model.ShowNewTaskStatusRequest) *ShowNewTaskStatusInvoker {
	requestDef := GenReqDefForShowNewTaskStatus()
	return &ShowNewTaskStatusInvoker{invoker.NewBaseInvoker(c.HcClient, request, requestDef)}
}
