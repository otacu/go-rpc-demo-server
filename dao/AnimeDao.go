package dao

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	proto "go-rpc-demo-server/proto"
	"log"
)

const (
	dataSource = "root:123456@tcp(127.0.0.1:3306)/python?charset=utf8"
)

type AnimeDao struct {
}

func (a *AnimeDao) QueryByAnimeId(animeId int32, res *proto.QueryAnimeInfoResponse) error {
	db, err0 := a.OpenDB()
	if err0 != nil {
		return err0
	}
	stmt, err1 := db.Prepare("select anime_id, en_name from tb_myanimelist_anime where anime_id=?")
	defer stmt.Close()
	if err1 != nil {
		return errors.New("prepare error")
	}
	row, err1 := stmt.Query(animeId)
	row.Next()
	err2 := row.Scan(&res.AnimeId, &res.EnName)

	if err2 != nil {
		return errors.New("query error")
	}
	return nil
}

func (a *AnimeDao) OpenDB() (db *sql.DB, err error) {
	db, err1 := sql.Open("mysql", dataSource)
	if err1 != nil {
		return db, errors.New("open database fail")
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	if err2 := db.Ping(); err2 != nil {
		return db, errors.New("ping database fail")
	}
	log.Println("connect success")
	return db, nil
}
