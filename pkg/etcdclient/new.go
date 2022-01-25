package etcdclient

import (
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func New() (*clientv3.Client, error) {
	logrus.Println("etcd client init")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{types.Server.ETCDHost + ":" + types.Server.ETCDPort},
		DialTimeout: types.Server.Timeout,
	})
	return cli, err
}