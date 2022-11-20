package test

import (
	"github.com/rcsrn/rmp/pkg/miner"
	"github.com/rcsrn/rmp/pkg/database"
	"testing"
	"os"
)

var testMiner *miner.Miner
var testRolas *database.Rola
var numberOfFiles int

func TestCreateNewMiner(t *testing.T) {
	workingDirectory, err := os.Getwd()
	testPath := workingDirectory + "/TestRolas"
	if err != nil {
		t.Errorf("This should not happen.")
	}
	
	f, err := os.Open(testPath)
	if err != nil {
		t.Errorf(err.Error())
	}
	
	files, err := f.Readdir(0)
	if err != nil {
		t.Errorf(err.Error())
	}
	
	numberOfFiles = len(files)
	
	testMiner = miner.CreateNewMiner(testPath)
}

func TestTraverse(t *testing.T) {
	testMiner.Traverse()
	filePaths := testMiner.GetFilePaths()
	if filePaths == nil {
		t.Errorf("Miner does not traverse correctly. The filePaths is empty.")
	} else if len(filePaths) != numberOfFiles {
		t.Errorf("The length of filePaths should be '%v' but '%v' has been gotten.",
			numberOfFiles, len(filePaths))
	}
}

