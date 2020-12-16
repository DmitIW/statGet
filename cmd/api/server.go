package api

import (
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
)

type StatisticServer struct {
	Addr   string
	Server *http.Server
	Mplx   *http.ServeMux
}

func (s *StatisticServer) Start() error {
	if err := s.Server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *StatisticServer) Stop() {
	if err := s.Server.Close(); err != nil {
		log.Printf("Error on closing server: %v\n", err)
	}
}

func GetMux(tConnection *tarantool.Connection) *http.ServeMux {
	mplx := http.NewServeMux()
	BindPWApi(tConnection, "src2dst_at", mplx, "/s2da")
	BindPWApi(tConnection, "src2dst_lg", mplx, "/s2dl")
	BindPWApi(tConnection, "dst2src_at", mplx, "/d2sa")
	BindPWApi(tConnection, "dst2src_lg", mplx, "/d2sl")
	BindPWApi(tConnection, "dst2proto_at", mplx, "/d2pa")
	BindPWApi(tConnection, "dst2proto_lg", mplx, "/d2pl")
	return mplx
}

func GetServer(addr string, tConnection *tarantool.Connection) StatisticServer {
	mplx := GetMux(tConnection)
	srv := &http.Server{
		Addr:    addr,
		Handler: mplx,
	}
	log.Printf("Server is created on %v\n", addr)
	return StatisticServer{
		Addr:   addr,
		Server: srv,
		Mplx:   mplx,
	}
}
