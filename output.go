package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type sumStruct struct {
	Sample      string
	ReadsNumber int `json:"Reads Number"`
	Bases       int `json:"Bases(bp)"`
	MeanLength  int `json:"Mean Length(bp)"`
	Longest     int `json:"Longest(bp)"`
	N50         int `json:"N50(bp)"`
	LenArray    []int
}

func outputJson(outStruct *sumStruct) {
	// 创建文件
	filePtr, err := os.Create(*outPut + "/" + *sampleName + "_stat.json")
	checkErr("文件创建失败", err)
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(outStruct)
	checkErr("Json编码错误", err)
}

func outputTxt(outStruct *sumStruct) {
	// 创建文件
	filePtr, err := os.Create(*outPut + "/" + *sampleName + "_stat.txt")
	defer filePtr.Close()
	checkErr("文件创建失败", err)
	sumContent := fmt.Sprintf("%s\t%d\t%d\t%d\t%d\t%d\n",
		outStruct.Sample,
		outStruct.ReadsNumber,
		outStruct.Bases,
		outStruct.MeanLength,
		outStruct.Longest,
		outStruct.N50,
	)
	_, err = filePtr.WriteString(
		"Sample\tReads number\tBases(bp)\tMean Length(bp)\tLongest(bp)\tN50(bp)\n" +
			sumContent,
	)
	checkErr("写入文件失败", err)
}
