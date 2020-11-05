package admin

import (
	`strconv`

	`github.com/class100/yunke-core`
	`github.com/storezhang/gox`
)

// GetOrg 取得机构的信息
func (ca *ClientAdmin) GetOrg() (org *yunke.Org, err error) {
	org = new(yunke.Org)
	err = ca.request(yunke.AdminApiGetOrgUrl, gox.HttpMethodGet, nil, map[string]string{
		"id": strconv.FormatInt(ca.Id, 10),
	}, org)

	return
}
