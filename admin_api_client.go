package admin

import (
	`strconv`

	`github.com/class100/yunke-core`
	`github.com/storezhang/gox`
)

// GetLastVersion 取得某个客户端类型的最新版本
func (ca *ClientAdmin) GetLastVersion(clientType core.ClientType) (client *core.BaseClient, err error) {
	client = new(core.BaseClient)
	err = ca.request(core.AdminApiGetFinalClientByTypeUrl, gox.HttpMethodGet, nil, map[string]string{
		"clientType": strconv.Itoa(int(clientType)),
	}, client)

	return
}
