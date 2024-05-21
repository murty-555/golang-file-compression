package main

import (
    "fmt"
    "bytes"
    "compress/gzip"
    "compress/flate"
    "io"
    "github.com/dsnet/compress/bzip2" 
	//the above has to be imported by using 
	//"go get github.com/dsnet/compress/bzip2" in terminal
	//but before that we have to create a module like
	// go mod init github/username/projectname
  )


//Creating Compressor interface with methods Compress and Decompress
type Compressor interface{
  Compress(data []byte)([]byte,error)
  Decompress(data []byte)([]byte,error)
}

//Creating gzip struct type
type GzipCompressor struct{}

//implementing gzip compress method
func(g GzipCompressor)Compress(data []byte)([]byte,error){
  var buf bytes.Buffer
  writer := gzip.NewWriter(&buf)
  _,err := writer.Write(data)
  errorCheck(err)
  err = writer.Close()
  errorCheck(err)
  return buf.Bytes(),nil
}

//implementing gzip decompress method
func(g GzipCompressor)Decompress(data []byte)([]byte,error){
  reader,err := gzip.NewReader(bytes.NewReader(data))
  errorCheck(err)
  defer reader.Close()
  return io.ReadAll(reader)
}

//creating defalte compressor struct type
type DeflateCompressor struct{}

//implementing Deflate compressor method
func(d DeflateCompressor)Compress(data []byte)([]byte,error){
  var buf bytes.Buffer
  writer,err := flate.NewWriter(&buf,flate.DefaultCompression)
  errorCheck(err)
  _,err = writer.Write(data)
  errorCheck(err)
  err = writer.Close()
  errorCheck(err)
  return buf.Bytes(),nil
}

//implementing Deflate Decompress method
func(d DeflateCompressor)Decompress(data []byte)([]byte,error){
  reader := flate.NewReader(bytes.NewReader(data))
  defer reader.Close()
  return io.ReadAll(reader)
}

//creating bzip2 compressor
type Bzip2Compressor struct{}

func (b Bzip2Compressor)Compress(data []byte)([]byte,error){
  var buf bytes.Buffer
  writer,err := bzip2.NewWriter(&buf,nil)
  errorCheck(err)
  _,err = writer.Write(data)
  errorCheck(err)
  err = writer.Close()
  errorCheck(err)
  return buf.Bytes(),nil
  
}

func(b Bzip2Compressor)Decompress(data []byte)([]byte,error){
  reader,err := bzip2.NewReader(bytes.NewReader(data),nil)
  errorCheck(err)
  defer reader.Close()
  return io.ReadAll(reader)
}

//code for reusablity for errors
func errorCheck(err error)([]byte,error){
	if err != nil{
		return nil,err
	}
	return nil,nil
}

//code for resuability for log errors
func errorLog(err error, name string){
	if err != nil{
		fmt.Println(name," error: ",err)
		return
	  }
}

  
func main(){
  
  data := []byte("Hello World!")
  
  //using gzipCompressor
  gzipCompressor := GzipCompressor{}
  
  gzipCompressedData,err := gzipCompressor.Compress(data)
  errorLog(err,"gzipCompression")
  
  fmt.Println("gzipCompressedData: ",gzipCompressedData)
  
  gzipDecompressedData,err := gzipCompressor.Decompress(gzipCompressedData)
  errorLog(err,"gzipDecompression")

  fmt.Println("gzipDecompressedData: ",string(gzipDecompressedData))
  
  //---------------------------------------------------
  //using DeflateCompressor
  deflateCompressor := DeflateCompressor{}
  
  deflateCompressedData,err := deflateCompressor.Compress(data)
  errorLog(err,"deflateCompression")

  fmt.Println("deflateCompressedData: ",deflateCompressedData)
  
  deflateDecompressedData,err := deflateCompressor.Decompress(deflateCompressedData)
  errorLog(err,"deflateDecompression")

  fmt.Println("deflateDecompressedData: ",string(deflateDecompressedData))
  
  //----------------------------------------------------
  //using Bzip2Compressor
  bzip2Compressor := Bzip2Compressor{}
  
  bzip2CompressedData,err := bzip2Compressor.Compress(data)
  errorLog(err,"bzip2Compression")
  fmt.Println("bzip2CompressedData: ",bzip2CompressedData)
  
  bzip2DecompressedData,err := bzip2Compressor.Decompress(bzip2CompressedData)
  errorLog(err,"bzip2Decompression")
  fmt.Println("bzip2DecompressedData: ",string(bzip2DecompressedData))
  
  
}