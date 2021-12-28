package main

import (
	"github.com/biogo/hts/bam"
	"github.com/biogo/hts/bgzf"
	"io"
	"log"
	"os"
)

func fileReader(file *string) (b *bam.Reader) {
	var r io.Reader
	if *file == "-" {
		r = os.Stdin
	} else {
		f, err := os.Open(*file)
		checkErr("could not open file", err)
		//defer f.Close()
		ok, err := bgzf.HasEOF(f)
		checkErr("could not open file", err)
		if !ok {
			log.Printf("file %q has no bgzf magic block: may be truncated", *file)
		}
		r = f
	}
	var err error
	b, err = bam.NewReader(r, *conc)
	checkErr("could not read bam", err)

	return b
}

func bamReader(b *bam.Reader, lenArray *[]int) {
	for {
		rec, err := b.Read()
		if err == io.EOF {
			break
		}
		checkErr("error reading bam:", err)
		*lenArray = append(*lenArray, len(rec.Seq.Expand()))
		//fmt.Printf("%s\n%v\n%v\n%d\n%s:%d..%d\n(%d)\n%d\n%s:%d\n%d\n%s\n%v\n%v\n------------\n",
		//	rec.Name,
		//	rec.Flags,
		//	rec.Cigar,
		//	rec.MapQ,
		//	rec.Ref.Name(),
		//	rec.Pos,
		//	rec.End(),
		//	rec.Bin(),
		//	rec.End()-rec.Pos,
		//	rec.MateRef.Name(),
		//	rec.MatePos,
		//	rec.TempLen,
		//	rec.Seq.Expand(),
		//	rec.Qual,
		//	rec.AuxFields)
	}
}
