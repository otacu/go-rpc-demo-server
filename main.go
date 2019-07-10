package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/micro/go-micro"
	"go-rpc-demo-server/dao"
	proto "go-rpc-demo-server/proto"
	"log"
)

const (
	dataSource = "root:123456@tcp(127.0.0.1:3306)/python?charset=utf8"
)

// 数据库连接池
var DB *sql.DB

type AnimeService struct {
}

func (s *AnimeService) QueryAnimeInfo(ctx context.Context, req *proto.QueryAnimeInfoRequest, res *proto.QueryAnimeInfoResponse) error {
	var animeDao = dao.AnimeDao{DB}
	err := animeDao.QueryByAnimeId(req.AnimeId, res)
	if err != nil {
		return err
	}
	return nil
}

// go run main.go
func main() {
	// 这里的open是创建数据库连接池，并不是打开连接，查询时才建立连接
	db, err := OpenDB()
	if err != nil {
		log.Println(err)
		return
	}
	DB = db

	service := micro.NewService(
		micro.Name("AnimeService"),
		micro.Version("latest"),
	)
	service.Init()
	// 服务默认注册在mdns，可以改成其它注册中心
	proto.RegisterAnimeServiceHandler(service.Server(), new(AnimeService))

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func OpenDB() (db *sql.DB, err error) {
	db, err1 := sql.Open("mysql", dataSource)
	if err1 != nil {
		return nil, errors.New("open database fail")
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	if err2 := db.Ping(); err2 != nil {
		return nil, errors.New("ping database fail")
	}
	log.Println("connect success")
	return db, nil
}
