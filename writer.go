package epubwriter

import (
    "archive/zip"
    "bytes"
    "fmt"
    "io/ioutil"
    "path"
    "strings"
    "text/template"
)

type Writer struct {
    Author string
    Chapters []*Chapter
    Cover *Cover
    Identifier string
    Language string
    Manifest []*Manifest
    Spine []*Spine
    Styles []*Style
    Title string
}

type Chapter struct {
    Bytes []byte
    Filename string
    Id string
    Title string
}

type Cover struct {
    Bytes []byte
    Filename string
}

type Manifest struct {
    Href string
    Id string
    MediaType string
    Properties string
}

type Spine struct {
    Idref string
    Linear bool
}

type Style struct {
    Bytes []byte
    Filename string
    Id string
}

func NewWriter(title string, author string, identifier string) *Writer {
    w := new(Writer)

    w.Author = author
    w.Chapters = make([]*Chapter, 0)
    w.Identifier = identifier
    w.Language = "en-US"
    w.Manifest = make([]*Manifest, 0)
    w.Spine = make([]*Spine, 0)
    w.Styles = make([]*Style, 0)
    w.Title = title

    w.appendManifest(
        "toc.ncx",
        "ncx",
        "application/x-dtbncx+xml",
        "")

    w.appendManifest(
        "toc.xhtml",
        "toc",
        "application/xhtml+xml",
        "nav")

    w.appendManifest(
        "titlepage.xhtml",
        "titlepage",
        "application/xhtml+xml",
        "")

    w.appendManifest(
        "cover.xhtml",
        "cover",
        "application/xhtml+xml",
        "")

    w.appendSpine("cover", false)
    w.appendSpine("titlepage", true)

    w.AppendStyle([]byte(stylesheet))

    return w
}

func (w *Writer) AppendChapter(title string, filename string, b []byte) *Chapter {
    c := new(Chapter)

    c.Bytes = b
    c.Filename = filename
    c.Title = title

    c.Id = path.Base(c.Filename)
    c.Id = strings.TrimSuffix(c.Id, path.Ext(c.Id))

    w.Chapters = append(w.Chapters, c)

    w.appendManifest(
        c.Filename,
        c.Id,
        "application/xhtml+xml",
        "")

    w.appendSpine(c.Id, true)

    return c
}

func (w *Writer) AppendStyle(b []byte) *Style {
    s := new(Style)

    s.Bytes = b
    s.Id = fmt.Sprintf("css%04d", len(w.Styles) + 1) 
    s.Filename = s.Id + ".css"

    w.Styles = append(w.Styles, s)

    w.appendManifest(
        s.Filename,
        s.Id,
        "text/css",
        "")
    
    return s
}

func (w *Writer) AddCover(filename string, b []byte) *Cover {
    c := new(Cover)

    c.Filename = filename
    c.Bytes = b

    w.Cover = c
    
    w.appendManifest(
        c.Filename,
        "cover-image",
        "image/jpeg",
        "cover-image")

    return c
}

func (w *Writer) Bytes() ([]byte, error) {
    w.appendSpine("toc", false)

    buf := new(bytes.Buffer)

    z := zip.NewWriter(buf)

    err := w.writeFile(z, "mimetype", []byte(mimetype))
    if err != nil {
        return nil, err
    }

    err = w.writeFile(z, "META-INF/container.xml", []byte(containerXml))
    if err != nil {
        return nil, err
    }
    
    err = w.writeManifest(z)
    if err != nil {
        return nil, err
    }

    err = w.writeNcx(z)
    if err != nil {
        return nil, err
    }

    err = w.writeTableOfContents(z)
    if err != nil {
        return nil, err
    }

    err = w.writeTitlePage(z)
    if err != nil {
        return nil, err
    }

    err = w.writeCover(z)
    if err != nil {
        return nil, err
    }

    if w.Cover != nil {
        err = w.writeFile(z, w.Cover.Filename, w.Cover.Bytes)
        if err != nil {
            return nil, err
        }
    }

    for _, c := range w.Chapters {
        err = w.writeChapter(z, c)
        if err != nil {
            return nil, err
        }
    }

    for _, s := range w.Styles {
        err = w.writeFile(z, s.Filename, s.Bytes)
        if err != nil {
            return nil, err
        }
    }

    z.Close()

    return buf.Bytes(), nil
}

func (w *Writer) WriteFile(filename string) error {
    b, err := w.Bytes()
    if err != nil {
        return err
    }

    return ioutil.WriteFile(filename, b, 0644)
}

func (w *Writer) appendManifest(href string, id string, mediaType string, properties string) *Manifest {
    m := new(Manifest)

    m.Href = href
    m.Id = id
    m.MediaType = mediaType
    m.Properties = properties

    w.Manifest = append(w.Manifest, m)

    return m
}

func (w *Writer) appendSpine(idref string, linear bool) *Spine {
    s := new(Spine)

    s.Idref = idref
    s.Linear = linear

    w.Spine = append(w.Spine, s)

    return s
}

func (w *Writer) writeChapter(z *zip.Writer, c *Chapter) error {
    t, err := template.New("T").Parse(chapter)
    if err != nil {
        return err
    }

    m := map[string] interface{} {
        "Html": string(c.Bytes),
        "Styles": w.Styles,
        "Title": c.Title,
    }

    buf := new(bytes.Buffer)

    err = t.ExecuteTemplate(buf, "T", m)
    if err != nil {
        return err
    }

    return w.writeFile(z, c.Filename, buf.Bytes())
}

func (w *Writer) writeCover(z *zip.Writer) error {
    t, err := template.New("T").Parse(cover)
    if err != nil {
        return err
    }

    m := map[string] interface{} {
        "Styles": w.Styles,
        "Title": w.Title,
    }

    if w.Cover != nil {
        m["Cover"] = w.Cover.Filename
    } else {
        m["Cover"] = coverimage
    }

    buf := new(bytes.Buffer)

    err = t.ExecuteTemplate(buf, "T", m)
    if err != nil {
        return err
    }

    return w.writeFile(z, "cover.xhtml", buf.Bytes())
}

func (w *Writer) writeFile(z *zip.Writer, s string, b []byte) error {
    f, err := z.Create(s)
    if err != nil {
        return err
    }

    _, err = f.Write(b)

    return err
}

func (w *Writer) writeManifest(z *zip.Writer) error {
    t, err := template.New("T").Parse(containerOpf)
    if err != nil {
        return err
    }

    buf := new(bytes.Buffer)

    err = t.ExecuteTemplate(buf, "T", w)
    if err != nil {
        return err
    }

    return w.writeFile(z, "container.opf", buf.Bytes())
}

func (w *Writer) writeNcx(z *zip.Writer) error {
    f := template.FuncMap {
        "add": func(a int, b int) int { return a + b },
    }

    t, err := template.New("T").Funcs(f).Parse(ncx)
    if err != nil {
        return err
    }

    buf := new(bytes.Buffer)

    err = t.ExecuteTemplate(buf, "T", w)
    if err != nil {
        return err
    }

    return w.writeFile(z, "toc.ncx", buf.Bytes())
}

func (w *Writer) writeTableOfContents(z *zip.Writer) error {
    t, err := template.New("T").Parse(tableOfContents)
    if err != nil {
        return err
    }

    buf := new(bytes.Buffer)

    err = t.ExecuteTemplate(buf, "T", w)
    if err != nil {
        return err
    }

    return w.writeFile(z, "toc.xhtml", buf.Bytes())
}

func (w *Writer) writeTitlePage(z *zip.Writer) error {
    t, err := template.New("T").Parse(titlePage)
    if err != nil {
        return err
    }

    m := map[string] interface{} {
        "Styles": w.Styles,
        "Title": w.Title,
    }

    if w.Cover != nil {
        m["Cover"] = w.Cover.Filename
    } else {
        m["Cover"] = coverimage
    }

    buf := new(bytes.Buffer)

    err = t.ExecuteTemplate(buf, "T", m)
    if err != nil {
        return err
    }

    return w.writeFile(z, "titlepage.xhtml", buf.Bytes())
}

