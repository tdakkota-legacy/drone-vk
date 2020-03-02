package plugin

import (
	"errors"
	"fmt"
	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/api/params"
	"github.com/urfave/cli"
)

type Plugin struct {
	c   *cli.Context
	api *api.VK
}

func (p Plugin) App(c *cli.Context) error {
	p.c = c

	token := c.String("token")
	if len(token) == 0 {
		return errors.New("invalid token")
	}

	p.api = api.Init(token)

	b, err := p.buildMessage()
	if err != nil {
		return err
	}

	_, err = p.api.MessagesSend(b)
	if err != nil {
		return err
	}

	return nil
}

func (p Plugin) buildMessage() (api.Params, error) {
	template := p.c.String("template")
	if template == "" {
		template = DroneTelegramTemplate
	}

	b := params.NewMessagesSendBuilder()

	info := ParseInfo(p.c)
	message, err := ExecuteTemplate(template, info)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %v", err)
	}
	b.Message(message)

	peerID := p.c.Int("peer_id")
	if peerID == 0 {
		return nil, errors.New("peer_id arg must be set")
	}
	b.PeerID(peerID)

	if stickerID := p.c.Int("sticker_id"); stickerID != 0 {
		b.StickerID(stickerID)
	}

	var attachments []string
	if attachment, err := p.checkImage(peerID); err != nil {
		return nil, err
	} else {
		attachments = append(attachments, attachment)
	}

	if attachment, err := p.checkDoc(peerID); err != nil {
		return nil, err
	} else {
		attachments = append(attachments, attachment)
	}

	b.Attachment(attachments)

	b.DontParseLinks(p.c.Bool("dont_parse_links"))
	b.RandomID(0)
	return b.Params, nil
}

func attachmentString(t string, userID, ID int, accessKey string) string {
	if accessKey != "" {
		return fmt.Sprintf("%s%d_%d_%s", t, userID, ID, accessKey)
	}

	return fmt.Sprintf("%s%d_%d", t, userID, ID)
}

func (p Plugin) checkImage(peerId int) (string, error) {
	if image := p.c.String("image"); image != "" {
		b := params.NewPhotosGetMessagesUploadServerBuilder()
		b.PeerID(peerId)

		r, err := p.api.PhotosGetMessagesUploadServer(b.Params)
		if err != nil {
			return "", fmt.Errorf("failed to get upload server: %v", err)
		}

		upload, err := NewUploader(r.UploadURL).UploadPhoto(image)
		if err != nil {
			return "", fmt.Errorf("failed to upload photo: %v", err)
		}

		b2 := params.NewPhotosSaveMessagesPhotoBuilder()
		b2.Photo(upload.Photo)
		b2.Server(upload.Server)
		b2.Hash(upload.Hash)

		save, err := p.api.PhotosSaveMessagesPhoto(b2.Params)
		if err != nil {
			return "", fmt.Errorf("failed to save photo: %v", err)
		}

		if len(save) < 1 {
			return "", fmt.Errorf("incorrect response: at least one object expected")
		}

		return attachmentString("photo", save[0].OwnerID, save[0].ID, save[0].AccessKey), err
	}
	return "", nil
}

func (p Plugin) checkDoc(peerId int) (string, error) {
	if doc := p.c.String("file.name"); doc != "" {
		b := params.NewDocsGetMessagesUploadServerBuilder()
		b.PeerID(peerId)

		if fileType := p.c.String("file.type"); fileType != "" {
			b.Type(fileType)
		}

		r, err := p.api.DocsGetMessagesUploadServer(b.Params)
		if err != nil {
			return "", fmt.Errorf("failed to get upload server: %v", err)
		}

		upload, err := NewUploader(r.UploadURL).UploadDoc(doc)
		if err != nil {
			return "", fmt.Errorf("failed to upload doc: %v", err)
		}

		b2 := params.NewDocsSaveBuilder()
		b2.File(upload.File)

		save, err := p.api.DocsSave(b2.Params)
		if err != nil {
			return "", fmt.Errorf("failed to save doc: %v", err)
		}

		return attachmentString("doc", save.Doc.OwnerID, save.Doc.ID, save.Doc.AccessKey), err
	}
	return "", nil
}
