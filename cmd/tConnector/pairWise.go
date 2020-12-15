package tConnector

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"gopkg.in/vmihailenco/msgpack.v2"
	"log"
)

type PairwiseStatistic struct {
	AprioriElement     uint16
	ProbabilityElement uint16
	Counter            uint32
	Timestamp          uint64
}

func (ps *PairwiseStatistic) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(2); err != nil {
		return err
	}

	if err := e.EncodeUint16(ps.AprioriElement); err != nil {
		return err
	}

	if err := e.EncodeUint16(ps.ProbabilityElement); err != nil {
		return err
	}

	if err := e.EncodeUint32(ps.Counter); err != nil {
		return err
	}

	if err := e.EncodeUint64(ps.Timestamp); err != nil {
		return err
	}

	return nil
}

func (ps *PairwiseStatistic) DecodeMsgpack(d *msgpack.Decoder) error {
	var (
		err error
		l   int
	)
	if l, err = d.DecodeArrayLen(); err != nil {
		return err
	}

	if l != 4 {
		return fmt.Errorf("PairwiseStatistic:Decode:: array len doesn't match: %d", l)
	}

	if ps.AprioriElement, err = d.DecodeUint16(); err != nil {
		return err
	}

	if ps.ProbabilityElement, err = d.DecodeUint16(); err != nil {
		return nil
	}

	if ps.Counter, err = d.DecodeUint32(); err != nil {
		return err
	}

	if ps.Timestamp, err = d.DecodeUint64(); err != nil {
		return err
	}

	return nil
}

type PairWiseTuple struct {
	Statistics PairwiseStatistic
}

func (t *PairWiseTuple) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := t.Statistics.EncodeMsgpack(e); err != nil {
		return err
	}
	return nil
}

func (t *PairWiseTuple) DecodeMsgpack(d *msgpack.Decoder) error {
	var (
		err error
	)
	if err = t.Statistics.DecodeMsgpack(d); err != nil {
		return err
	}
	return nil
}

type PWDConnection struct {
	conn TConnection
	pwt  []PairWiseTuple
}

func PairWiseAgent(connection *tarantool.Connection, spaceName string) PWDConnection {
	return PWDConnection{
		conn: TarantoolAgent(connection, spaceName),
	}
}

func (pwdC *PWDConnection) SelectCounter(aprioriElement uint16, probabilityElement uint16) uint32 {
	if err := pwdC.conn.Select([]interface{}{aprioriElement, probabilityElement}, &pwdC.pwt); err != nil {
		log.Printf("WARNINIG:PWDConnection:SelectCounter:: %v\n", err)
		return 0
	}
	if len(pwdC.pwt) == 0 {
		return 0
	}
	return pwdC.pwt[0].Statistics.Counter
}

func (pwdC *PWDConnection) SelectTotal(aprioriElement uint16) uint32 {
	return pwdC.SelectCounter(aprioriElement, 0)
}

func (pwdC *PWDConnection) SelectMean(aprioriElement uint16) uint32 {
	return pwdC.SelectCounter(0, aprioriElement)
}
