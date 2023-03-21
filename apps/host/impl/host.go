package impl

import (
	"context"
	"github.com/elisaovivo/restful-api-demo/apps/host"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	return nil, nil
}
func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}
func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.QueryHostRequest) (*host.Host, error) {
	return nil, nil
}
func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.QueryHostRequest) (*host.Host, error) {
	return nil, nil
}
func (i *HostServiceImpl) DeleteHost(context.Context, *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
