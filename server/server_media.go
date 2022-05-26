package server

import (
	"database/sql"
	"go.uber.org/zap"
	"net/url"

	"github.com/status-im/status-go/ipfs"
)

type MediaServer struct {
	Server

	db         *sql.DB
	downloader *ipfs.Downloader
}

// NewMediaServer returns a *MediaServer
func NewMediaServer(db *sql.DB, downloader *ipfs.Downloader) (*MediaServer, error) {
	err := generateTLSCert()
	if err != nil {
		return nil, err
	}

	s := &MediaServer{
		Server:     NewServer(globalCertificate, localhost),
		db:         db,
		downloader: downloader,
	}

	return s, nil
}

func (s *MediaServer) withMediaHandlers() {
	l := s.logger.Named("withMediaHandlers").With(zap.String("name", "withMediaHandlers"))
	l.Info("has been fired")
	s.WithHandlers(HandlerPatternMap{
		imagesPath:     handleImage(s.db, s.logger),
		audioPath:      handleAudio(s.db, s.logger),
		identiconsPath: handleIdenticon(s.logger),
		ipfsPath:       handleIPFS(s.downloader, s.logger),
	})
}

func (s *MediaServer) Start() error {
	l := s.logger.Named("MediaServer.Start()").With(zap.String("name", "MediaServer.Start()"))
	l.Info("has been fired")
	s.withMediaHandlers()
	go s.listenAndServe()
	return nil
}

func (s *MediaServer) MakeImageServerURL() string {
	u := s.MakeBaseURL()
	u.Path = basePath + "/"
	return u.String()
}

func (s *MediaServer) MakeIdenticonURL(from string) string {
	u := s.MakeBaseURL()
	u.Path = identiconsPath
	u.RawQuery = url.Values{"publicKey": {from}}.Encode()

	return u.String()
}

func (s *MediaServer) MakeImageURL(id string) string {
	u := s.MakeBaseURL()
	u.Path = imagesPath
	u.RawQuery = url.Values{"messageId": {id}}.Encode()

	return u.String()
}

func (s *MediaServer) MakeAudioURL(id string) string {
	u := s.MakeBaseURL()
	u.Path = audioPath
	u.RawQuery = url.Values{"messageId": {id}}.Encode()

	return u.String()
}

func (s *MediaServer) MakeStickerURL(stickerHash string) string {
	u := s.MakeBaseURL()
	u.Path = ipfsPath
	u.RawQuery = url.Values{"hash": {stickerHash}}.Encode()

	return u.String()
}
