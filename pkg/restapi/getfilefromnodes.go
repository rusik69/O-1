package restapi

import (
	"errors"

	"github.com/loqutus/O-1/pkg/client"
	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func getFileFromNodes(fileName string, nodes []string) error {
	downloaded := false
	for _, node := range nodes {
		logrus.Println("Trying node", node)
		types.Client.HostName = node
		types.Client.Port = types.Server.ListenPort
		err := client.Download(fileName)
		if err != nil {
			continue
		} else {
			downloaded = true
			logrus.Println("Found file", fileName, "at node", node)
		}
	}
	if !downloaded {
		return errors.New("download failed")
	} else {
		return nil
	}
}