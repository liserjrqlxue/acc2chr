package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/liserjrqlxue/goUtil/osUtil"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
	"github.com/liserjrqlxue/version"
)

var (
	acc2chr = map[string]string{
		"NC_000001.10": "1",
		"NC_000002.11": "2",
		"NC_000003.11": "3",
		"NC_000004.11": "4",
		"NC_000005.9":  "5",
		"NC_000006.11": "6",
		"NC_000007.13": "7",
		"NC_000008.10": "8",
		"NC_000009.11": "9",
		"NC_000010.10": "10",
		"NC_000011.9":  "11",
		"NC_000012.11": "12",
		"NC_000013.10": "13",
		"NC_000014.8":  "14",
		"NC_000015.9":  "15",
		"NC_000016.9":  "16",
		"NC_000017.10": "17",
		"NC_000018.9":  "18",
		"NC_000019.9":  "19",
		"NC_000020.10": "20",
		"NC_000021.8":  "21",
		"NC_000022.10": "22",
		"NC_000023.10": "X",
		"NC_000024.9":  "Y",
		"NC_012920.1":  "MT",
	}
	sharp = regexp.MustCompile(`^#`)
)

var (
	vcf = flag.String(
		"v",
		"",
		"input vcf",
	)
)

func main() {
	version.LogVersion()
	flag.Parse()

	var scanner *bufio.Scanner
	if *vcf == "" || *vcf == "-" {
		var f = os.Stdin
		defer simpleUtil.DeferClose(f)
		scanner = bufio.NewScanner(f)
	} else if strings.HasSuffix(*vcf, "gz") || strings.HasSuffix(*vcf, "bgz") {
		var f = osUtil.Open(*vcf)
		defer simpleUtil.DeferClose(f)
		var g = simpleUtil.HandleError(gzip.NewReader(f)).(*gzip.Reader)
		defer simpleUtil.DeferClose(g)
		scanner = bufio.NewScanner(g)
	} else {
		var f = osUtil.Open(*vcf)
		scanner = bufio.NewScanner(f)
	}

	for scanner.Scan() {
		if line := scanner.Text(); sharp.MatchString(line) {
			fmt.Println(line)
		} else {
			var array = strings.Split(line, "\t")
			var chr, ok = acc2chr[array[0]]
			if ok {
				array[0] = chr
			}
			fmt.Println(strings.Join(array, "\t"))
		}
	}
	simpleUtil.CheckErr(scanner.Err())
}
