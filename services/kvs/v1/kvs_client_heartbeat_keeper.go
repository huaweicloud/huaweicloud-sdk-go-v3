package v1

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kvs/v1/model"
	"strings"
	"sync"
	"time"
)

func StartKvsClientHeartbeatKeeper(kvsClientMap *sync.Map, heartbeatInterval int, heartbeatThreadMaxCount int, ctx context.Context) {
	ticker := time.NewTicker(time.Duration(heartbeatInterval) * time.Second)
	semaphore := make(chan struct{}, heartbeatThreadMaxCount)
	defer ticker.Stop()
	defer close(semaphore)

	for range ticker.C {
		select {
		case <-ctx.Done():
			return
		default:
			kvsClientMap.Range(func(_, value interface{}) bool {
				managedClient, ok := value.(*KvsManagedClient)
				if ok {
					semaphore <- struct{}{}
					go kvsClientHeartbeatCheck(managedClient, semaphore)
				}
				return true
			})
		}

	}
}

func kvsClientHeartbeatCheck(managedClient *KvsManagedClient, semaphore chan struct{}) {
	defer func() {
		<-semaphore
	}()
	request := &model.CheckHealthRequest{}
	request.Body = &model.CheckHealthRequestBody{}
	_, err := managedClient.KvsClient.CheckHealth(request)

	if err == nil || strings.Contains(err.Error(), "400") {
		managedClient.IsUsable.Store(true)
	} else {
		managedClient.IsUsable.Store(false)
	}
}
