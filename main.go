package main

import (
	"context"
	"github.com/micro/go-micro"
	"go-rpc-demo-server/dao"
	proto "go-rpc-demo-server/proto"
	"log"
)

type AnimeService struct {
}

func (s *AnimeService) QueryAnimeInfo(ctx context.Context, req *proto.QueryAnimeInfoRequest, res *proto.QueryAnimeInfoResponse) error {
	animeDao := dao.AnimeDao{}
	err := animeDao.QueryByAnimeId(req.AnimeId, res)
	if err != nil {
		return err
	}
	return nil
}

// go run main.go
func main() {
	service := micro.NewService(
		micro.Name("AnimeService"),
		micro.Version("latest"),
	)
	service.Init()
	proto.RegisterAnimeServiceHandler(service.Server(), new(AnimeService))

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
