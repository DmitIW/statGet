package api

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"log"
	"net/http"
	"statGet/cmd/randomDist"
	"statGet/cmd/tConnector"
	"statGet/cmd/utility"
)

type PairWiseAPI struct {
	conn *tConnector.PWDConnection
}

func calc(mean, std, lambda float64) float64 {
	var (
		counter float64
	)
	counter = randomDist.ABSNormal(mean, std)
	return randomDist.Poisson(counter, lambda)
}

func (pwa *PairWiseAPI) Calc(aprioriElement uint16, probabilityElement uint16) float64 {
	var (
		counter, total, mean uint32
	)

	if total = pwa.conn.SelectTotal(aprioriElement); total == 0 {
		return calc(0.0, 1.0, 17.0)
	}

	if counter = pwa.conn.SelectCounter(aprioriElement, probabilityElement); counter == 0 {
		mean = pwa.conn.SelectMean(aprioriElement)
		return calc(0, 1.5*float64(mean), float64(mean))
	}

	return float64(counter) / float64(total)
}

func (pw *PairWiseAPI) PairWiseHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			a, p        uint16
			probability float64
			err         error
		)
		a = utility.GetUint16(r.FormValue("a"))
		p = utility.GetUint16(r.FormValue("p"))

		probability = pw.Calc(a, p)

		if _, err = fmt.Fprintf(w, "%f", probability); err != nil {
			log.Printf("Error on sending probability: %v\n", err)
		}
	})
}

func (pw *PairWiseAPI) SpaceSizeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			size uint32
			err  error
		)
		size = pw.conn.Size()
		if _, err = fmt.Fprintf(w, "%d", size); err != nil {
			log.Printf("Error on sending size: %v\n", err)
		}
	})
}

func (pwa *PairWiseAPI) Bind(mplx *http.ServeMux, route string) {
	probabilityRoute := route + "_prob"
	sizeRoute := route + "_size"
	mplx.Handle(probabilityRoute, pwa.PairWiseHandler())
	mplx.Handle(sizeRoute, pwa.SpaceSizeHandler())
}

func createPWApi(conn *tarantool.Connection, spaceName string) PairWiseAPI {
	connector := tConnector.PairWiseAgent(conn, spaceName)
	return PairWiseAPI{
		conn: &connector,
	}
}

func BindPWApi(conn *tarantool.Connection, spaceName string, mplx *http.ServeMux, route string) {
	pwa := createPWApi(conn, spaceName)
	pwa.Bind(mplx, route)
}
