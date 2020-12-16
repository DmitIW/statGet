package tConnector

import (
	"github.com/tarantool/go-tarantool"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type TConnection struct {
	connection *tarantool.Connection

	spaceName string
	spaceID   uint32
	indexID   uint32
}

func TarantoolAgent(connection *tarantool.Connection, spaceName string) TConnection {
	space := connection.Schema.Spaces[spaceName]
	spaceID := space.Id
	indexID := space.Indexes["primary"].Id
	return TConnection{
		connection: connection,
		spaceName:  spaceName,
		spaceID:    spaceID,
		indexID:    indexID,
	}
}

func (c *TConnection) Select(key []interface{}, result interface{}) error {
	if err := c.connection.SelectTyped(c.spaceID, c.indexID, 0, 1, tarantool.IterEq,
		key, result); err != nil {
		return err
	}
	return nil
}

type SizeTuple struct {
	Size uint
}

func (t *SizeTuple) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeArrayLen(1); err != nil {
		return err
	}

	if err := e.EncodeUint(t.Size); err != nil {
		return err
	}
	return nil
}

func (t *SizeTuple) DecodeMsgpack(d *msgpack.Decoder) error {
	var (
		err error
		//l   int
	)
	//
	//if l, err = d.DecodeArrayLen(); err != nil {
	//	return err
	//}
	//
	//if l != 1 {
	//	return fmt.Errorf("SizeTuple:Decode:: array len doesn't match: %d", l)
	//}

	if t.Size, err = d.DecodeUint(); err != nil {
		return err
	}
	return nil
}

func (c *TConnection) Size() (uint, error) {
	var (
		err  error
		size []SizeTuple
	)
	if err = c.connection.EvalTyped("return box.space."+c.spaceName+":count()", []interface{}{}, &size); err != nil {
		return 0, err
	}
	if len(size) == 0 {
		return 0, nil
	}
	return size[0].Size, nil
}
