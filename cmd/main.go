package main

import (
	"fmt"

	"github.com/ediallocyf/output/pkg/cli"
	//"github.com/ediallocyf/output/pkg/util"
)

// // --------------------------------------
// const (
// 	directory          = "../assets/"
// 	fileExtention      = ".txt"
// 	chLength           = 8
// 	firstCh       byte = ' '
// 	lastCh        byte = '~'
// )

// //----------------------------------------

// func getFinalString() {

// 	//----------------------------------------
// 	patternContent, errPattern := util.ReadFileToStr(directory, fileExtention )

// 	if errPattern != nil {
// 		fmt.Println("Error : loading pattern file failed")
// 		return
// 	}
// 	//--------------------------------
// 	patternMap, errPatternMap := util.ContentToMap(patternContent, chLength)

// 	if errPatternMap != nil {
// 		fmt.Println(errPatternMap)
// 		return
// 	}
// 	//--------------------------------------------
// 	cliStr, errClistr := cli.ReadCli()

// 	if errClistr != nil {
// 		//fmt.Println(errClistr)
// 		return
// 	}
// 	//----------------------------------------------

// 	finalStr := util.MapToStr(cliStr, patternMap, chLength)
// 	fmt.Println(finalStr)
// 	//-----------------------------------------
// }

func main() {
	r, _ := cli.OutputFlag()
	fmt.Println(r)
}