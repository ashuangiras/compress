package compress

import (
	"bytes"
	
	"io"
	"io/ioutil"
	"log"
	lzma "github.com/itchio/lzma"
)



func LZMACompress(payload []byte) ([]byte, error) {

	pr, pw := io.Pipe()
	//defer pr.Close()
	//defer pw.Close()
	go func() {
	defer pw.Close()
	var z io.WriteCloser
		// var inFile *os.File
		// var err error

	r := bytes.NewReader(payload)

	z = lzma.NewWriter(pw)
	defer z.Close()
	_, err := io.Copy(z, r)
	if err != nil {
		log.Fatal(err.Error())
	}
	// log.Println("Number of Bytes Compressed written : ",nWritten)
}()
	
	// write into outFile from pr
	defer pr.Close()
	
	compressed, err := ioutil.ReadAll(pr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return compressed, err
}


func LZMADecompress(payload []byte) ([]byte, error) {
	pr, pw := io.Pipe()

	go func() {
	r := bytes.NewReader(payload)
		defer pw.Close()
		
	_, err := io.Copy(pw, r)
	if err != nil {
		log.Fatal(err.Error())
	}

	// log.Println("Number of Byte Decompressed written : ",nWritten)
	}()
	
	// write into outFile from z
	defer pr.Close()
	z := lzma.NewReader(pr)
	defer z.Close()
	
	output, err := ioutil.ReadAll(z)
	if err != nil {
		log.Fatal(err.Error())
	}

	return output, err
}
