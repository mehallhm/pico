package pager

import (
	"encoding/binary"

	"github.com/mehallhm/picosql/compiler"
)

func SerializeRow(row *compiler.InsertionRow) [291]byte {
	var out []byte = make([]byte, 4, 291)

	// var username [32]byte
	// copy(username[:], args[2])
	//
	// var email [255]byte
	// copy(email[:], args[3])

	out = binary.LittleEndian.AppendUint32(out, row.Id)
	out = append(out, row.Username[:]...)
	out = append(out, row.Email[:]...)

	return [291]byte(out)
}

func DeserializeRow(row [291]byte) *compiler.InsertionRow {
	id := binary.LittleEndian.Uint32(row[0:4])

	// username := string(row[4:36])
	// email := string(row[36:291])

	return &compiler.InsertionRow{
		Id:       id,
		Username: [32]byte(row[4:36]),
		Email:    [255]byte(row[36:291]),
	}
}
