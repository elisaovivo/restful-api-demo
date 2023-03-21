package impl

import (
	"context"
	"github.com/elisaovivo/restful-api-demo/apps/host"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	if err := ins.Validate(); err != nil {
		return nil, err
	}
	ins.InjectDefault()

	if err := i.save(ctx, ins); err != nil {
		return nil, err
	}
	return ins, nil
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
