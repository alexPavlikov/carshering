package config

import (
	"encoding/json"
	"io/fs"
	"os"

	"github.com/alexPavlikov/scud.ru/algorithms"
	"github.com/alexPavlikov/scud.ru/models"
)

func Config() {
	file, err := os.Open("./config/config.cfg")
	algorithms.CheckErr(err)
	defer file.Close()

	var stat fs.FileInfo

	stat, err = file.Stat()
	algorithms.CheckErr(err)

	readByte := make([]byte, stat.Size())

	_, err = file.Read(readByte)
	algorithms.CheckErr(err)

	err = json.Unmarshal(readByte, &models.Config)
	algorithms.CheckErr(err)
}
