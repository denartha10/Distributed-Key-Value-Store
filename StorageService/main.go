package main

import "errors"

type StorageNode struct {
	ID        string
	IPADDRESS string
	STORAGE   map[byte]string
}

func (s StorageNode) Get(key byte) (string, error) {
	if value, exists := s.STORAGE[key]; exists {
		return value, nil
	}
	return "", errors.New("Key not found")
}

func (s StorageNode) Put(key byte, value string) {
	s.STORAGE[key] = value
}

func (s StorageNode) Pop(key byte) (string, error) {
	if value, exists := s.STORAGE[key]; exists {
		delete(s.STORAGE, key)
		return value, nil
	}
	return "", errors.New("Key doesent exist")
}
