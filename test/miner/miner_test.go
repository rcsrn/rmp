package test

import (
	"github.com/rcsrn/rmp/pkg/miner"
	"github.com/rcsrn/rmp/pkg/database"
	"testing"
	"os"
	"fmt"
)

var testMiner *miner.Miner
var testRolas *database.Rola

func TestCreateNewMiner(t *testing.T) {
	workingDirectory, err := os.Getwd()
	testPath := workingDirectory + "/TestRolas"
	if err != nil {
		fmt.Println("This should not happen.")
	}
	testMiner = miner.CreateNewMiner(testPath)
}

func TestTraverse(t *testing.T) {
	testMiner.Traverse()
	if filePaths := testMiner.GetFilePaths(); filePaths == nil || len(filePaths) != 3 {
		fmt.Println("miner does not traverses correctly.")
	}
}

// func TestGetRolas(t *testing.T) {
// 	testMiner.MineTags()
// }

// func createTestRolas() {
// 	testRolas := make([]*database.Rola, 0)
// 	testRolas = append(testRolas, database.CreateNewRola())
// 	testRolas = append(testRolas, database.CreateNewRola())
// 	testRolas = append(testRolas, database.CreateNewRola())

// 	testRolas[0].SetPerformer("Meshuggah")
// }
