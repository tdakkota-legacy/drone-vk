package plugin

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type UploaderOption func(u *Uploader)

func WithClient(client *http.Client) UploaderOption {
	return func(u *Uploader) {
		u.client = client
	}
}

type Uploader struct {
	client    *http.Client
	serverUrl string
}

func NewUploader(serverUrl string, options ...UploaderOption) *Uploader {
	u := &Uploader{serverUrl: serverUrl}

	for _, option := range options {
		option(u)
	}

	if u.client == nil {
		u.client = http.DefaultClient
	}

	return u
}

func (u *Uploader) uploadData(filename, field string, file io.Reader, i interface{}) error {
	body := new(bytes.Buffer)
	m := multipart.NewWriter(body)

	part, err := m.CreateFormFile(field, filename)
	if err != nil {
		return err
	}

	if _, err = io.Copy(part, file); err != nil {
		return err
	}

	if err = m.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", u.serverUrl, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", m.FormDataContentType())

	resp, err := u.client.Do(req)
	if err != nil {
		return err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	return json.NewDecoder(resp.Body).Decode(i)
}

type PhotoUploadResponse struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

func (u *Uploader) UploadPhoto(filename string) (response PhotoUploadResponse, err error) {
	f, err := os.Open(filename)

	if err != nil {
		return
	}
	defer f.Close()

	err = u.uploadData(filename, "photo", f, &response)
	return
}

type DocUploadResponse struct {
	File string `json:"file"`
}

func (u *Uploader) UploadDoc(filename string) (response DocUploadResponse, err error) {
	f, err := os.Open(filename)

	if err != nil {
		return
	}
	defer f.Close()

	err = u.uploadData(filename, "file", f, &response)
	return
}
