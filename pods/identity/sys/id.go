package sys

import (
	"os"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

func GenerateID() (id int64, err error) {

	nodeID := int64(1)

	if nodev, exists := os.LookupEnv("NODE"); exists {
		if noden, err := strconv.ParseInt(nodev, 10, 64); err != nil {
			nodeID = noden
		}
	}

	node, err := snowflake.NewNode(nodeID)

	if err != nil {
		return id, err
	}

	return node.Generate().Int64(), nil
}

func MustGenerateID() int64 {
	id, _ := GenerateID()
	return id
}
