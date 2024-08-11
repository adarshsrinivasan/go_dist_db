package main

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type PathTransformFunc func(string) string

var DefaultPathTransformFunc = func(key string) string {return key}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}


type Store struct {
	StoreOpts
}

func NewStorage(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) writeStream(key string, reader io.Reader)error {
	pathName := s.PathTransformFunc(key)

	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		return err
	}

	fileName := "someFileName"
	pathAndFileName := pathName + "/" + fileName

	filePtr, err := os.Create(pathAndFileName)
	if err != nil {
		return err
	}

	n, err := io.Copy(filePtr, reader)
	if err != nil {
		return err
	}

	log.Infof("Wrote %d bytes to disk at: %s", n, pathAndFileName)

	return nil
}