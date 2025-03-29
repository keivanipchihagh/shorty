package kgs

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/keivanipchihagh/shorty/internal/services/base62"
)

func GenerateId() (int64, string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0, "", err
	}

	id := node.Generate()
	idInt := id.Int64()
	return idInt, base62.Encode(idInt), nil
}
