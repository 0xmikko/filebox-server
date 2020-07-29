/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package repository

import (
	"github.com/MikaelLazarev/filebox-server/config"
	"github.com/MikaelLazarev/filebox-server/core"
	ipfs "github.com/ipfs/go-ipfs-api"
	"io"
)

type IPFSClient struct {
	shell *ipfs.Shell
}

func NewIPFSClient(config *config.Config) core.IPFSRepositoryI {
	sh := ipfs.NewShell(config.IpfsEndpoint)
	return &IPFSClient{shell: sh}
}

func (sh *IPFSClient) AddFile(r io.Reader) (string, error) {
	return sh.shell.Add(r)
}

func (sh *IPFSClient) GetFile(ipfsHash string) ([]byte, error) {
	panic("implement me")
}
