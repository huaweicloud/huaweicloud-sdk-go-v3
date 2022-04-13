package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/antiddos/v1/model"
)

type AntiDDoSClient struct {
	HcClient *http_client.HcHttpClient
}

func NewAntiDDoSClient(hcClient *http_client.HcHttpClient) *AntiDDoSClient {
	return &AntiDDoSClient{HcClient: hcClient}
}

func AntiDDoSClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//配置用户的默认防护策略。配置防护策略后，新购买的资源在自动开启防护时，会按照该默认防护策略进行配置。
func (c *AntiDDoSClient) CreateDefaultConfig(request *model.CreateDefaultConfigRequest) (*model.CreateDefaultConfigResponse, error) {
	requestDef := GenReqDefForCreateDefaultConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDefaultConfigResponse), nil
	}
}

//删除用户配置的默认防护策略。
func (c *AntiDDoSClient) DeleteDefaultConfig(request *model.DeleteDefaultConfigRequest) (*model.DeleteDefaultConfigResponse, error) {
	requestDef := GenReqDefForDeleteDefaultConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDefaultConfigResponse), nil
	}
}

//查询用户配置信息，用户可以通过此接口查询是否接收某类告警，同时可以配置是手机短信还是电子邮件接收告警信息。
func (c *AntiDDoSClient) ShowAlertConfig(request *model.ShowAlertConfigRequest) (*model.ShowAlertConfigResponse, error) {
	requestDef := GenReqDefForShowAlertConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowAlertConfigResponse), nil
	}
}

//查询用户配置的默认防护策略。
func (c *AntiDDoSClient) ShowDefaultConfig(request *model.ShowDefaultConfigRequest) (*model.ShowDefaultConfigResponse, error) {
	requestDef := GenReqDefForShowDefaultConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDefaultConfigResponse), nil
	}
}

//更新用户配置信息，用户可以通过此接口更新是否接收某类告警，同时可以配置是手机短信还是电子邮件接收告警信息。
func (c *AntiDDoSClient) UpdateAlertConfig(request *model.UpdateAlertConfigRequest) (*model.UpdateAlertConfigResponse, error) {
	requestDef := GenReqDefForUpdateAlertConfig()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateAlertConfigResponse), nil
	}
}

//查询用户所有EIP的Anti-DDoS防护状态信息，用户的EIP无论是否绑定到云服务器，都可以进行查询。
func (c *AntiDDoSClient) ListDDosStatus(request *model.ListDDosStatusRequest) (*model.ListDDosStatusResponse, error) {
	requestDef := GenReqDefForListDDosStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDDosStatusResponse), nil
	}
}

//查询指定EIP在过去24小时之内的异常事件信息，异常事件包括清洗事件和黑洞事件，查询延迟在5分钟之内。
func (c *AntiDDoSClient) ListDailyLog(request *model.ListDailyLogRequest) (*model.ListDailyLogResponse, error) {
	requestDef := GenReqDefForListDailyLog()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDailyLogResponse), nil
	}
}

//查询指定EIP在过去24小时之内的防护流量信息，流量的间隔时间单位为5分钟。
func (c *AntiDDoSClient) ListDailyReport(request *model.ListDailyReportRequest) (*model.ListDailyReportResponse, error) {
	requestDef := GenReqDefForListDailyReport()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDailyReportResponse), nil
	}
}

//查询系统支持的Anti-DDoS防护策略配置的可选范围，用户根据范围列表选择适合自已业务的防护策略进行Anti-DDoS流量清洗。
func (c *AntiDDoSClient) ListNewConfigs(request *model.ListNewConfigsRequest) (*model.ListNewConfigsResponse, error) {
	requestDef := GenReqDefForListNewConfigs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNewConfigsResponse), nil
	}
}

//查询用户所有Anti-DDoS防护周统计情况，包括一周内DDoS拦截次数和攻击次数、以及按照被攻击次数进行的排名信息等统计数据。系统支持当前时间之前四周的周统计数据查询，超过这个时间的请求是查询不到统计数据的。
func (c *AntiDDoSClient) ListWeeklyReports(request *model.ListWeeklyReportsRequest) (*model.ListWeeklyReportsResponse, error) {
	requestDef := GenReqDefForListWeeklyReports()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListWeeklyReportsResponse), nil
	}
}

//查询配置的Anti-DDoS防护策略，用户可以查询指定EIP的Anti-DDoS防护策略。
func (c *AntiDDoSClient) ShowDDos(request *model.ShowDDosRequest) (*model.ShowDDosResponse, error) {
	requestDef := GenReqDefForShowDDos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDDosResponse), nil
	}
}

//查询指定EIP的Anti-DDoS防护状态。
func (c *AntiDDoSClient) ShowDDosStatus(request *model.ShowDDosStatusRequest) (*model.ShowDDosStatusResponse, error) {
	requestDef := GenReqDefForShowDDosStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDDosStatusResponse), nil
	}
}

//用户查询指定的Anti-DDoS防护配置任务，得到任务当前执行的状态。
func (c *AntiDDoSClient) ShowNewTaskStatus(request *model.ShowNewTaskStatusRequest) (*model.ShowNewTaskStatusResponse, error) {
	requestDef := GenReqDefForShowNewTaskStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowNewTaskStatusResponse), nil
	}
}

//更新指定EIP的Anti-DDoS防护策略配置。调用成功，只是说明服务节点收到了关闭更新配置请求，操作是否成功需要通过任务查询接口查询该任务的执行状态，具体请参考查询Anti-DDoS任务。
func (c *AntiDDoSClient) UpdateDDos(request *model.UpdateDDosRequest) (*model.UpdateDDosResponse, error) {
	requestDef := GenReqDefForUpdateDDos()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateDDosResponse), nil
	}
}
