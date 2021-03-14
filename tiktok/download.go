package tiktok

import (
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func DownloadTikTok(url string, percent *int) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	*percent = 10

	fileName := GenerateRandName()
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()
	*percent = 15

	_, err = io.Copy(out, res.Body)
	*percent = 100
	return err
}

func GenerateRandName() string {
	randName := "tiktok-"

	rand.Seed(time.Now().UnixNano())
	ints := rand.Perm(5)

	for _, v := range ints {
		randName += strconv.Itoa(v)
	}

	return randName + ".mp4"
}
