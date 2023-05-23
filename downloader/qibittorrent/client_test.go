package qibittorrent_test

import (
	"autobangumi-go/downloader/qibittorrent"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/anacrolix/torrent/bencode"
	torrent "github.com/anacrolix/torrent/metainfo"
	"github.com/stretchr/testify/require"
)

func TestQbittorent(t *testing.T) {
	qb, err := qibittorrent.NewQbittorrentClient("http://nas.lyqingye.com:8888", "admin", "adminadmin", "/downloads")
	require.NoError(t, err)
	require.NotNil(t, qb)
	err = qb.Login()
	require.NoError(t, err)

	bz, err := ioutil.ReadFile("/home/lyqingye/Downloads/4f0379c68a3078f5ed8504be57e0b827cfe5b812.torrent")
	require.NoError(t, err)
	torrent, err := torrent.Load(bytes.NewBuffer(bz))
	require.NoError(t, err)
	info, err := torrent.UnmarshalInfo()
	require.NoError(t, err)
	//"361bfee217db001fceb309f0bebd8b53745c92fb"
	properties, err := qb.GetTorrentProperties(torrent.HashInfoBytes().HexString())
	require.Equal(t, err, qibittorrent.ErrTorrentNotFound)
	require.Nil(t, properties)

	// change filename
	info.Name = "test.mp4"
	infoBytes, err := bencode.Marshal(&info)
	require.NoError(t, err)
	torrent.InfoBytes = infoBytes

	newTorrent := bytes.Buffer{}
	err = torrent.Write(&newTorrent)
	require.NoError(t, err)

	_, err = qb.AddTorrent("", newTorrent.Bytes(), "/downloads")
	require.NoError(t, err)
	err = qb.Logout()
	require.NoError(t, err)
}

func TestListTorrent(t *testing.T) {
	qb, err := qibittorrent.NewQbittorrentClient("http://nas.lyqingye.com:8888", "admin", "adminadmin", "/downloads")
	require.NoError(t, err)
	require.NotNil(t, qb)
	err = qb.Login()
	require.NoError(t, err)
	list, err := qb.ListAllTorrent(qibittorrent.FilterAllTorrentList)
	for _, torrent := range list {
		t.Log(torrent.Hash)
	}
	require.NoError(t, err)
}

func TestGetTorrentContent(t *testing.T) {
	qb, err := qibittorrent.NewQbittorrentClient("http://localhost:8080", "admin", "adminadmin", "/downloads")
	require.NoError(t, err)
	require.NotNil(t, qb)
	err = qb.Login()
	require.NoError(t, err)
	bz, err := ioutil.ReadFile("/home/lyqingye/Downloads/a55ed28af3d95bf54f74c0abe4ca0ebedbbac347.torrent")
	require.NoError(t, err)
	hash, err := qb.AddTorrent("", bz, "")
	require.NoError(t, err)
	properties, err := qb.GetTorrentProperties(hash)
	require.NoError(t, err)
	require.NotNil(t, properties)
	files, err := qb.GetTorrentContent(hash, []int64{})
	require.NoError(t, err)
	require.NotNil(t, files)
	for _, fi := range files {
		err = qb.RenameFile(hash, fi.Name, "renamed")
		require.NoError(t, err)
	}

}
