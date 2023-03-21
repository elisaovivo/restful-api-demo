package impl

import (
	"context"
	"fmt"
	"github.com/elisaovivo/restful-api-demo/apps/host"
)

// 完成对象和数据直接的转换

func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {

	//两次插入操作，开启事务
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error,%s", err)
	}

	//处理事务提交方式
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				fmt.Errorf("rollback error,%s", err)
			}
			return
		}
		if err := tx.Commit(); err != nil {
			fmt.Errorf("commit error,%s", err)
		}
	}()

	rstmt, err := tx.Prepare(InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()
	_, err = rstmt.Exec(ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP)
	if err != nil {
		return err
	}
	dstmt, err := tx.Prepare(InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()
	_, err = dstmt.Exec(ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber)
	if err != nil {
		return err
	}
	return nil
}
