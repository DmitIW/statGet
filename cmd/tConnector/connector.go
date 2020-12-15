package tConnector

import (
	"github.com/tarantool/go-tarantool"
)

type TConnection struct {
	connection *tarantool.Connection

	spaceID uint32
	indexID uint32
}

func TarantoolAgent(connection *tarantool.Connection, spaceName string) TConnection {
	space := connection.Schema.Spaces[spaceName]
	spaceID := space.Id
	indexID := space.Indexes["primary"].Id
	return TConnection{
		connection: connection,
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
