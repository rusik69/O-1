package restapi

import (
	"crypto/sha256"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/loqutus/O-1/pkg/types"
	"github.com/sirupsen/logrus"
)

func PostFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[1:]
	fileNameWithPath := types.Server.LocalDir + "/" + fileName
	logrus.Println("PostFile " + fileName)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(err, w)
		return
	}

	err = os.WriteFile(fileNameWithPath, body, 0644)
	if err != nil {
		Error(err, w)
		return
	}
	fi, err := os.Stat(fileNameWithPath)
	if err != nil {
		Error(err, w)
		return
	}
	fileSize := fi.Size()
	hash := sha256.New()
	if _, err := io.Copy(hash, r.Body); err != nil {
		Error(err, w)
		return
	}
	SHA256 := hash.Sum(nil)
	fileInfo := types.FileInfo{
		Name:   fileName,
		Size:   fileSize,
		SHA256: string(SHA256),
		Nodes:  []string{},
	}
	fileInfoJSON, _ := json.Marshal(fileInfo)
	_, err = types.Server.Cli.Put(*types.Server.Ctx, fileName, string(fileInfoJSON))
	if err != nil {
		Error(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}
