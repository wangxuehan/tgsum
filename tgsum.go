package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

var (
	//file       = flag.String("bam", "", "Input bam (empty for stdin)")
	sampleName = flag.String("sample", "sample", "Sample name")
	outPut     = flag.String("output", ".", "Output directory")
	conc       = flag.Int("threads", 0, "Number of threads to use (0 = auto)")
	help       = flag.Bool("help", false, "Display help")
)

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: tgsum [-bam bamName] [-sample sample] [-output outputDir] [-threads threads]
       统计三代下机数据bam的长度信息.

Options:
`)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, `
Example: tgsum -bam reads.bam    # 输出到当前目录sample_stat.{json,txt}
         tgsum -bam reads.bam -sample wangxuehan -output result    # 输出到result目录下wangxuehan_stat.{json,txt}
         tgsum -bam reads1.bam -bam reads2.bam    # 支持多个bam合并统计, 适用于一个样本多个数据

Result: 
        header:Sample    Reads Number    Bases(bp)    Mean Length(bp)    Longest(bp)    N50(bp)

Author:     王雪涵
version:    v0.1
`)
}

type sliceFlag []string

func (f *sliceFlag) String() string {
	return fmt.Sprintf("%v", []string(*f))
}

func (f *sliceFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func main() {
	var file sliceFlag
	flag.Var(
		&file, "bam", "Input bam (empty for stdin), for example: -bam=a.bam -bam=b.bam",
	)
	flag.Usage = usage
	flag.Parse()
	if *help || len(file) < 1 {
		flag.Usage()
		os.Exit(0)
	}
	//设置协程数量
	num := *conc
	if num == 0 {
		num = runtime.NumCPU()
	}
	//循环读取所有bam
	var totalLength []int
	for _, bamFile := range file {
		if bamFile == "" {
			flag.Usage()
			os.Exit(0)
		}
		b := fileReader(&bamFile)
		defer b.Close()
		bamReader(b, &totalLength)
	}
	outStruct := sumInfo(totalLength)
	outputJson(outStruct)
	outputTxt(outStruct)
}
