package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

var (
	VideoFileTempBytes []byte = []byte("abcde12345hellobye")

	VideoFileTempChunkSize VideoFileByteCounter = 5

	VideoFileTempExpectedOutput [][]byte = [][]byte{
		[]byte("abcde"),
		[]byte("12345"),
		[]byte("hello"),
		[]byte("bye" + string(byte(0)) + string(byte(0))),
		[]byte(strings.Repeat(string(byte(0)), 5)),
	}

	VideoFileTempExpectedMinByte []VideoFileByteCounter = []VideoFileByteCounter{
		0, 5, 10, 15, 18,
	}

	VideoFileTempExpectedMaxByte []VideoFileByteCounter = []VideoFileByteCounter{
		5, 10, 15, 18, 18,
	}

	VideoFileTempExpectedIsFinished []bool = []bool{
		false, false, false, false, true,
	}
)

func TestGetVideoFileReader(t *testing.T) {
	/*
		Write a temporary file and write a known
		pattern of bytes to it, and check that the
		function will return the currect bytes.
	*/
	tmpFile, err := os.CreateTemp("", "*.mp4")
	if err != nil {
		t.Errorf("Error during CreateTemp: %+v\n", err)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	if _, err = tmpFile.Write(VideoFileTempBytes); err != nil {
		t.Errorf("Error during writing to temporary file: %+v\n", err)
	}

	_, err = GetVideoFileReader(tmpFile.Name(), VideoFileTempChunkSize)
	if err != nil {
		t.Errorf("Error during GetVideoFileReader(): %+v\n", err)
	}

}

func TestGetVideoFileReaderRead(t *testing.T) {
	/*
		Write a temporary file and write a known
		pattern of bytes to it, and check that the
		function will return the currect bytes.
	*/
	tmpFile, err := os.CreateTemp("", "*.mp4")
	if err != nil {
		t.Errorf("Error during CreateTemp: %+v\n", err)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	if _, err = tmpFile.Write(VideoFileTempBytes); err != nil {
		t.Errorf("Error during writing to temporary file: %+v\n", err)
	}

	vfr, err := GetVideoFileReader(tmpFile.Name(), VideoFileTempChunkSize)
	if err != nil {
		t.Logf("Error during GetVideoFileReader(): %+v\n", err)
		t.Skip()
	}

	for i := 0; i < len(VideoFileTempExpectedOutput); i++ {
		t.Logf("Iteration %d:\n", i)
		next, err := vfr.GetNextChunk()
		if err != nil {
			t.Errorf("Error on get next chunk: %+v\n", err)
		}

		t.Logf("Min bytes: %d, MaxByte %d, Length %d, isFinished %v, Bytes: %s\n", next.MinByte, next.MaxByte, next.Length, next.Finished, next.Bytes)

		/*
			Check that all values are as expected
		*/
		if !reflect.DeepEqual(VideoFileTempExpectedOutput[i], next.Bytes) {
			t.Error("Bytes output doesn't match")
		}
		if !reflect.DeepEqual(VideoFileTempExpectedMinByte[i], next.MinByte) {
			t.Error("Minimum byte doesn't match")
		}
		if !reflect.DeepEqual(VideoFileTempExpectedIsFinished[i], next.Finished) {
			t.Error("Finished bool doesn't match")
		}
		if !reflect.DeepEqual(VideoFileTempExpectedMinByte[i], next.MinByte) {
			t.Error("Minimum byte doesn't match")
		}
		if !reflect.DeepEqual(VideoFileTempExpectedMaxByte[i], next.MaxByte) {
			t.Error("Maximum byte doesn't match")
		}
	}
}
