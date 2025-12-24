package main

import (
	"flag"
	"strings"

	pqcgen "github.com/BondMachineHQ/pqcgen/pkg/pqcgen"
)

var verbose = flag.Bool("v", false, "Verbose")
var debug = flag.Bool("d", false, "Debug")

var qBits = flag.Int("qbits", 0, "Number of qbits (mandatory)")
var l = flag.Int("l", 0, "Number of blocks (mandatory)")
var outputType = flag.String("otype", "", "Output type: qiskitsymb, bmqsim (mandatory)")
var blockSeq = flag.String("bseq", "1", "Block sequence (mandatory)")

var listBlocks = flag.Bool("listblocks", false, "List available blocks and exit")

func init() {
	flag.Parse()
	if !*listBlocks {
		if *qBits <= 0 {
			panic("qbits must be > 0")
		}
		if *l <= 0 {
			panic("l must be > 0")
		}
		switch *outputType {
		case "qiskitsymb", "bmqsim":
		default:
			panic("otype must be one of: qiskitsymb, bmqsim")
		}
		blocks := strings.Split(*blockSeq, ",")
		if len(blocks) != *l {
			panic("number of blocks in bseq must be equal to l")
		}
		for _, b := range blocks {
			if !pqcgen.CheckBlockExists(b) {
				panic("block " + b + " does not exist")
			}
		}
	}
}

func main() {
	if *listBlocks {
		pqcgen.ListAvailableBlocks()
		return
	}

}
