package main

import (
	"flag"

	"github.com/BondMachineHQ/pqcgen/pkg/pqcgen"
)

var verbose = flag.Bool("v", false, "Verbose")
var debug = flag.Bool("d", false, "Debug")

var qBits = flag.Int("qbits", 4, "Number of qbits (mandatory)")
var l = flag.Int("l", 1, "Number of blocks (mandatory)")
var outputType = flag.String("otype", "qiskitsymb", "Output type: qiskitsymb, bmqsim (mandatory)")
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
	}
}

func main() {
	if *listBlocks {
		pqcgen.ListAvailableBlocks()
		return
	}
}
