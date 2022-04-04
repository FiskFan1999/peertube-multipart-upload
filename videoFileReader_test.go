package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

type VideoFileTestCase struct {
	Original      []byte
	ChunkSize     VideoFileByteCounter
	ExpChunks     [][]byte
	ExpMinByte    []VideoFileByteCounter
	ExpMaxByte    []VideoFileByteCounter
	ExpIsFinished []bool
}

var (
	NULL = string(byte(0))

	VideoFileTestCases []VideoFileTestCase = []VideoFileTestCase{
		{
			[]byte("abcde12345hellobye"),
			5,
			[][]byte{
				[]byte("abcde"),
				[]byte("12345"),
				[]byte("hello"),
				[]byte("bye" + string(byte(0)) + string(byte(0))),
				[]byte(strings.Repeat(string(byte(0)), 5)),
			},
			[]VideoFileByteCounter{
				0, 5, 10, 15, 18,
			},
			[]VideoFileByteCounter{
				4, 9, 14, 17, 18,
			},
			[]bool{
				false, false, false, false, true,
			},
		},
		{
			[]byte("depictedsedationsambainggoodbye"),
			8,
			[][]byte{
				[]byte("depicted"),
				[]byte("sedation"),
				[]byte("sambaing"),
				[]byte("goodbye" + NULL),
				[]byte(strings.Repeat(NULL, 8)),
			},
			[]VideoFileByteCounter{
				0, 8, 16, 24, 31,
			},
			[]VideoFileByteCounter{
				7, 15, 23, 30, 31,
			},
			[]bool{
				false, false, false, false, true,
			},
		},
	}
)

func TestGetVideoFileReader(t *testing.T) {
	/*
		Write a temporary file and write a known
		pattern of bytes to it, and check that the
		function will return the currect bytes.
	*/
	for i, c := range VideoFileTestCases {
		tmpFile, err := os.CreateTemp("", "*.mp4")
		if err != nil {
			t.Errorf("Error during CreateTemp: %+v\n", err)
		}
		defer tmpFile.Close()
		defer os.Remove(tmpFile.Name())

		if _, err = tmpFile.Write(c.Original); err != nil {
			t.Errorf("Case %d: Error during writing to temporary file: %+v\n", i, err)
		}

		_, err = GetVideoFileReader(tmpFile.Name(), c.ChunkSize)
		if err != nil {
			t.Errorf("Case %d: Error during GetVideoFileReader(): %+v\n", i, err)
		}
	}
}

func TestGetVideoFileReaderRead(t *testing.T) {
	/*
		Write a temporary file and write a known
		pattern of bytes to it, and check that the
		function will return the currect bytes.
	*/
	for j, c := range VideoFileTestCases {
		tmpFile, err := os.CreateTemp("", "*.mp4")
		if err != nil {
			t.Errorf("Case %d: Error during CreateTemp: %+v\n", j, err)
		}
		defer tmpFile.Close()
		defer os.Remove(tmpFile.Name())

		if _, err = tmpFile.Write(c.Original); err != nil {
			t.Errorf("Case %d: Error during writing to temporary file: %+v\n", j, err)
		}

		vfr, err := GetVideoFileReader(tmpFile.Name(), c.ChunkSize)
		if err != nil {
			t.Logf("Case %d: Error during GetVideoFileReader(): %+v\n", j, err)
			t.Skip()
		}

		//for i := 0; i < len(c.ExpChunks); i++ {
		for i := range c.ExpChunks {
			t.Logf("Case %d Iteration %d:\n", j, i)
			next, err := vfr.GetNextChunk()
			if err != nil {
				t.Errorf("Case %d: Error on get next chunk: %+v\n", j, err)
			}

			t.Logf("Min bytes: %d, MaxByte %d, Length %d, isFinished %v, Bytes: %s\n", next.MinByte, next.MaxByte, next.Length, next.Finished, next.Bytes)

			/*
				Check that all values are as expected
			*/
			if !reflect.DeepEqual(c.ExpChunks[i], next.Bytes) {
				t.Error("Bytes output doesn't match")
			}
			if !reflect.DeepEqual(c.ExpMinByte[i], next.MinByte) {
				t.Error("Minimum byte doesn't match")
			}
			if !reflect.DeepEqual(c.ExpIsFinished[i], next.Finished) {
				t.Error("Finished bool doesn't match")
			}
			if !reflect.DeepEqual(c.ExpMinByte[i], next.MinByte) {
				t.Error("Minimum byte doesn't match")
			}
			if !reflect.DeepEqual(c.ExpMaxByte[i], next.MaxByte) {
				t.Error("Maximum byte doesn't match")
			}
		}
	}
}
