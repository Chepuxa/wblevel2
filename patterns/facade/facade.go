package main

/*
Реализовать паттерн фасад, объяснить применимость, плюсы и минусы, а также реальные примеры использования паттерна на практике.

Фасад — это структурный паттерн проектирования, который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку.

Применимость:
- Когда нужно представить простой или урезанный интерфейс к сложной подсистеме.
- Когда нужно разложить подсистему на отдельные слои.

Плюсы:
- Изолирует клиентов от компонентов сложной подсистемы.

Минусы:
- Сложность поддерживания фасада при росте сложности системы.
- Фасад рискует стать божественным объектом, привязанным ко всем классам программы.
*/

type VideoFile struct {
	VideoBytes []byte
	Format string
}

type BitrateReader struct {
	quality int
}

func (br *BitrateReader) readBitrate(vf *VideoFile) int {
	return len(vf.VideoBytes) + br.quality
}

type CodecFactory struct {}

func (cf *CodecFactory) getCodec(format string, bitrate int) Codec {
	if format == "mp4" {
		return Codec{bitrate << 1}
	}
	if format == "ogg" {
		return Codec{bitrate << 2}
	}
	return Codec{bitrate}
}

type Codec struct {
	padding int
}

func (c *Codec) transform(video []byte, toFormat string) []byte {
	transformed := make([]byte, len(video))
	for _, v := range video {
		var newByte byte
		if toFormat == "mp4" {
			newByte = byte(c.padding) + v
		}
		if toFormat == 	"ogg" {
			newByte = byte(c.padding) - v
		}
		transformed = append(transformed, newByte)
	}
	return transformed
}

//Фасад, предоставляющий пользователю метод для преобразовывания видео из одного формата в другой
type VideoConverter struct {
	BitrateReader *BitrateReader
	CodecFactory *CodecFactory
}

func (vc *VideoConverter) convert(vf *VideoFile, toFormat string) VideoFile {
	bitrate := vc.BitrateReader.readBitrate(vf)
	codec := vc.CodecFactory.getCodec(vf.Format, bitrate)
	videoBytes := codec.transform(vf.VideoBytes, toFormat)
	return VideoFile{videoBytes, toFormat}
}

func newVideoConverter(q int) *VideoConverter {
	return &VideoConverter{
		&BitrateReader{q},
		&CodecFactory{},
	}
}

func main() {
	video := []byte{1, 2, 3, 4}
	vf := &VideoFile{video, "mp4"}
	vc := newVideoConverter(10)
	vc.convert(vf, "ogg")
}