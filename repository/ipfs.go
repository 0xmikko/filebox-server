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
	"strconv"
	"time"
)

type IPFSClient struct {
	shell  *ipfs.Shell
	tmpDir string
}

func NewIPFSClient(config *config.Config) core.IPFSRepositoryI {
	sh := ipfs.NewShell(config.IpfsEndpoint)
	return &IPFSClient{shell: sh, tmpDir: config.TemporaryDir}
}

func (sh *IPFSClient) AddFile(r io.Reader) (string, error) {
	return sh.shell.Add(r)
}

func (sh *IPFSClient) GetFile(ipfsHash string) (string, error) {
	tmpFile := sh.tmpDir + strconv.Itoa(int(time.Now().Unix()))
	if err := sh.shell.Get(ipfsHash, tmpFile); err != nil {
		return "", err
	}
	return tmpFile, nil
}
