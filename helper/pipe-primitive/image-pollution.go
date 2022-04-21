package pipe_primitive

import (
	"flag"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io/ioutil"
	"os"
)

func ImagePollution(primitive Primitive) (err error) {
	flagSet := flag.NewFlagSet(imagePollutionExpName(primitive), flag.ContinueOnError)
	var source, dest string
	flagSet.StringVar(&source, "source", "", "the path of file with evil content")
	flagSet.StringVar(&dest, "destination", "", "the path of file you want to pollution")
	awesome_error.CheckFatal(flagSet.Parse(os.Args[1:]))
	payload, err := ioutil.ReadFile(source)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return WriteImage(primitive, dest, payload)
}

func WriteImage(primitive Primitive, path string, payload []byte) (err error) {
	err = MultiWrite(primitive, path, payload)
	return
}
