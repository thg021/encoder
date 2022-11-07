package services

import (
	"context"
	"io/ioutil"
	"os"
	"os/exec"

	"log"

	"cloud.google.com/go/storage"
	"github.com/thg021/encoder/application/repositories"
	"github.com/thg021/encoder/domain"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	bucket := client.Bucket(bucketName)
	obj := bucket.Object(v.Video.FilePath)
	r, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}
	defer r.Close()

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	f, err := os.Create(os.Getenv("localStoragePath") + "/" + v.Video.ID + ".mp4")
	if err != nil {
		return err
	}

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	defer f.Close()

	log.Printf("Video %v has been stored", v.Video.ID)

	return nil
}

func (v *VideoService) Fragment() error {

	//vamos criar um diretorio e ter permissao
	err := os.Mkdir(os.Getenv("localStoragePath")+"/"+v.Video.ID, os.ModePerm)
	if err != nil {
		return err
	}

	//vamos pegar o caminho de onde esta nosso video
	source := os.Getenv("localStoragePath") + "/" + v.Video.ID + ".mp4"
	target := os.Getenv("localStoragePath") + "/" + v.Video.ID + ".frag"

	//este comando vai executar o servico do "bento4" para fragmentar o video
	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	printOutput(output)

	return nil
}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("=====> Output: %s\n", string(out))
	}
}
