package epubwriter

import (
    "archive/zip"
    "bytes"
    "io/ioutil"
)

type Writer struct {
    Author string
    Chapters []*Chapter
    Identifier string
    Title string
}

type Chapter struct {
    Bytes []byte
    Filename string
    Title string
}

func NewWriter(title string, author string, identifier string) *Writer {
    w := new(Writer)
    w.Author = author
    w.Chapters = make([]*Chapter, 0)
    w.Identifier = identifier
    w.Title = title
    return w
}

func (w *Writer) AppendChapter(title string, filename string, b []byte) *Chapter {
    c := new(Chapter)
    c.Bytes = b
    c.Filename = filename
    c.Title = title

    w.Chapters = append(w.Chapters, c)

    return c
}

func (w *Writer) WriteFile(filename string) error {
    buf := new(bytes.Buffer)

    z := zip.NewWriter(buf)

    for _, c := range w.Chapters {
        err := w.writeFile(z, c.Filename, c.Bytes)
        if err != nil {
            return err
        }
    }

    z.Close()

    return ioutil.WriteFile(filename, buf.Bytes(), 0644)
}

func (w *Writer) writeFile(z *zip.Writer, s string, b []byte) error {
    f, err := z.Create(s)
    if err != nil {
        return err
    }

    _, err = f.Write(b)

    return err
}
