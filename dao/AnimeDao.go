package dao

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	proto "go-rpc-demo-server/proto"
)

type AnimeDao struct {
	DB *sql.DB
}

func (a *AnimeDao) QueryByAnimeId(animeId int32, res *proto.QueryAnimeInfoResponse) error {
	stmt, err1 := a.DB.Prepare("select anime_id, en_name from tb_myanimelist_anime where anime_id=?")
	// 表示return后关闭连接
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

