package epubwriter

var chapter = `<?xml version="1.0" encoding="UTF-8"?>
<html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops">
<head>
<meta charset="utf-8"/>
<title>{{.Title}}</title>
{{range .Styles}}
<link rel="stylesheet" href="{{.Filename}}" type="text/css"/>
{{end}}
</head>
<body class="epub">
<section class="body-rw Chapter-rw" epub:type="bodymatter chapter">
{{.Html}}
</section>
</body>
</html>`

var containerOpf = `<?xml version="1.0" encoding="utf-8"?>
<package unique-identifier="pub-id" version="3.0" xml:lang="en" xmlns="http://www.idpf.org/2007/opf">
	<metadata xmlns:dc="http://purl.org/dc/elements/1.1/">
		<dc:title id="title">{{.Title}}</dc:title>
		<dc:creator id="creator">{{.Author}}</dc:creator>
		<dc:identifier id="pub-id">{{.Identifier}}</dc:identifier>
		<dc:language>en-US</dc:language>
	</metadata>
	<manifest>
        {{range .Manifest}}
		<item href="{{.Href}}" id="{{.Id}}" media-type="{{.MediaType}}"{{if .Properties}} properties="{{.Properties}}"{{end}}/>
        {{end}}
	</manifest>
	<spine toc="ncx">
        {{range .Spine}}
		<itemref idref="{{.Idref}}" linear="{{if .Linear}}yes{{else}}no{{end}}"/>
        {{end}}
	</spine>
</package>`

var containerXml = `<?xml version="1.0" encoding="utf-8"?>
<container version="1.0" xmlns="urn:oasis:names:tc:opendocument:xmlns:container">
	<rootfiles>
		<rootfile full-path="container.opf" media-type="application/oebps-package+xml"/>
	</rootfiles>
</container>`

var cover = `<?xml version="1.0" encoding="UTF-8"?>
<html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops">
<head>
<title>{{.Title}}</title>
{{range .Styles}}
<link rel="stylesheet" href="{{.Filename}}" type="text/css"/>
{{end}}
<meta charset="utf-8"/>
</head>
<body>
<div>
<img src="{{.Cover}}" alt="Cover Image" title="Cover Image"/>
</div>
</body>
</html>`

var coverimage = `data:image/jpeg;base64,
/9j/4AAQSkZJRgABAgAAZABkAAD/7AARRHVja3kAAQAEAAAAUAAA/+4ADkFkb2JlAGTAAAAAAf/b
AIQAAgICAgICAgICAgMCAgIDBAMCAgMEBQQEBAQEBQYFBQUFBQUGBgcHCAcHBgkJCgoJCQwMDAwM
DAwMDAwMDAwMDAEDAwMFBAUJBgYJDQsJCw0PDg4ODg8PDAwMDAwPDwwMDAwMDA8MDAwMDAwMDAwM
DAwMDAwMDAwMDAwMDAwMDAwM/8AAEQgB1AGbAwERAAIRAQMRAf/EAJIAAQEBAQEBAAAAAAAAAAAA
AAAJCAoGBwEBAAMBAQEBAQAAAAAAAAAAAAUGBwQIAgMBEAEAAAIHCQEBAQADAAAAAAAABQhyAwQG
B1cYAbECM5PTtTgJlBESMSITEQEAAAQACQ0BAQACAgMAAAAAAQIDBHHBctIzk9OUBlGxElITUzRU
BRU1BxcRITGBQaGRIjL/2gAMAwEAAhEDEQA/AL8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA83e2
9927iQG13ovdFqmBXfsFbZqm2xW0bOL/AMqritloq7LU/wC9vBw8W3h2cVbW8HDt4tuz+cP9/vFt
2cOzbt2P8/8AMYQw/wCQ/wC4x/yGGJ/Jowj0ZZpo/wAjH+SwjNNH+f7/ACWWWEZpo8kssIzRj/kI
Ri+Yampfs3rsfuq330Ze8p6yTOc3b1vKXm6XWxNTUv2b12P3VZ0Ze8p6yTOO3reUvN0utiampfs3
rsfuqzoy95T1kmcdvW8pebpdbE1NS/ZvXY/dVnRl7ynrJM47et5S83S62Jqal+zeux+6rOjL3lPW
SZx29byl5ul1sTU1L9m9dj91WdGXvKeskzjt63lLzdLrYmpqX7N67H7qs6MveU9ZJnHb1vKXm6XW
xNTUv2b12P3VZ0Ze8p6yTOO3reUvN0utiampfs3rsfuqzoy95T1kmcdvW8pebpdbE1NS/ZvXY/dV
nRl7ynrJM47et5S83S62Jqal+zeux+6rOjL3lPWSZx29byl5ul1sTU1L9m9dj91WdGXvKeskzjt6
3lLzdLrYmpqX7N67H7qs6MveU9ZJnHb1vKXm6XWxNTUv2b12P3VZ0Ze8p6yTOO3reUvN0utiampf
s3rsfuqzoy95T1kmcdvW8pebpdbE1NS/ZvXY/dVnRl7ynrJM47et5S83S62Jqal+zeux+6rOjL3l
PWSZx29byl5ul1sTU1L9m9dj91WdGXvKeskzjt63lLzdLrYmpqX7N67H7qs6MveU9ZJnHb1vKXm6
XWxNTUv2b12P3VZ0Ze8p6yTOO3reUvN0utiampfs3rsfuqzoy95T1kmcdvW8pebpdbE1NS/ZvXY/
dVnRl7ynrJM47et5S83S62Jqal+zeux+6rOjL3lPWSZx29byl5ul1sTU1L9m9dj91WdGXvKeskzj
t63lLzdLrYmpqX7N67H7qs6MveU9ZJnHb1vKXm6XWxNTUv2b12P3VZ0Ze8p6yTOO3reUvN0utiam
pfs3rsfuqzoy95T1kmcdvW8pebpdbE1NS/ZvXY/dVnRl7ynrJM47et5S83S62Jqal+zeux+6rOjL
3lPWSZx29byl5ul1sTU1L9m9dj91WdGXvKeskzjt63lLzdLrYmpqX7N67H7qs6MveU9ZJnHb1vKX
m6XWxNTUv2b12P3VZ0Ze8p6yTOO3reUvN0utiampfs3rsfuqzoy95T1kmcdvW8pebpdbE1NS/ZvX
Y/dVnRl7ynrJM47et5S83S62Jqal+zeux+6rOjL3lPWSZx29byl5ul1sTU1L9m9dj91WdGXvKesk
zjt63lLzdLrYmpqX7N67H7qs6MveU9ZJnHb1vKXm6XWxNTUv2b12P3VZ0Ze8p6yTOO3reUvN0uti
ampfs3rsfuqzoy95T1kmcdvW8pebpdbF7K5OKuHWJFpjFkuJe6wXqroBVWWujG2H8e2tq6jhtvFX
8NR/qs2bP8bdvFts1Z/12bdu3Z/P7t2bNnFw/wB+Y/z+/wAhNLHJjCaH/wAyxjD/AKfvTjPNL0pq
dSn/AL/P5Up1KM0f+P8AYS1ZZJow/wB//UIdH+/2H9/sI/z6C/j6AAAAAAAAAAZNnj9XcTacC87D
3Ne6GfBFN8N/JUMuHO5/FOehwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8X
WL0bRzYcTIPsbxdLIxxVfTDPAAAAAAAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAA
AAAAAABk2eP1dxNpwLzsPc17oZ8EU3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAVW+YnOxyoXY3xdYvRtHNhxMg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfB
FN8N/JUMuHO5/FOehwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzY
cTIPsbxdLIxxVfTDPAAAAAAAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAAB
k2eP1dxNpwLzsPc17oZ8EU3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
VW+YnOxyoXY3xdYvRtHNhxMg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JU
MuHO5/FOehwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbx
dLIxxVfTDPAAAAAAAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAABk2eP1dx
NpwLzsPc17oZ8EU3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVW+YnOx
yoXY3xdYvRtHNhxMg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JUMuHO5/F
OehwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbxdLIxxVf
TDPAAAAAAAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAABk2eP1dxNpwLzsP
c17oZ8EU3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVW+YnOxyoXY3xd
YvRtHNhxMg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JUMuHO5/FOehwAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbxdLIxxVfTDPAAAA
AAAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAABk2eP1dxNpwLzsPc17oZ8E
U3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVW+YnOxyoXY3xdYvRtHNh
xMg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JUMuHO5/FOehwAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbxdLIxxVfTDPAAAAAAAAAAG
TZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB
Vb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAABk2eP1dxNpwLzsPc17oZ8EU3w38lQ
y4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVW+YnOxyoXY3xdYvRtHNhxMg+xvF
0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JUMuHO5/FOehwAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbxdLIxxVfTDPAAAAAAAAAAGTZ4/V3E
2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABVb5ic7H
KhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAABk2eP1dxNpwLzsPc17oZ8EU3w38lQy4c7n8U
56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVW+YnOxyoXY3xdYvRtHNhxMg+xvF0sjHFV9
MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JUMuHO5/FOehwAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbxdLIxxVfTDPAAAAAAAAAAGTZ4/V3E2nAvOw9
zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABVb5ic7HKhdjfF1
i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAABk2eP1dxNpwLzsPc17oZ8EU3w38lQy4c7n8U56HAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVW+YnOxyoXY3xdYvRtHNhxMg+xvF0sjHFV9MM8AAAA
AAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JUMuHO5/FOehwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbxdLIxxVfTDPAAAAAAAAAAGTZ4/V3E2nAvOw9zXuhnwR
TfDfyVDLhzufxTnocAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABVb5ic7HKhdjfF1i9G0c2H
EyD7G8XSyMcVX0wzwAAAAAAAAABk2eP1dxNpwLzsPc17oZ8EU3w38lQy4c7n8U56HAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAVW+YnOxyoXY3xdYvRtHNhxMg+xvF0sjHFV9MM8AAAAAAAAAAZ
Nnj9XcTacC87D3Ne6GfBFN8N/JUMuHO5/FOehwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAF
VvmJzscqF2N8XWL0bRzYcTIPsbxdLIxxVfTDPAAAAAAAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVD
LhzufxTnocAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8X
SyMcVX0wzwAAAAAAAAABk2eP1dxNpwLzsPc17oZ8EU3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAVW+YnOxyoXY3xdYvRtHNhxMg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcT
acC87D3Ne6GfBFN8N/JUMuHO5/FOehwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzsc
qF2N8XWL0bRzYcTIPsbxdLIxxVfTDPAAAAAAAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxT
nocAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0
wzwAAAAAAAAABk2eP1dxNpwLzsPc17oZ8EU3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAVW+YnOxyoXY3xdYvRtHNhxMg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3
Ne6GfBFN8N/JUMuHO5/FOehwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XW
L0bRzYcTIPsbxdLIxxVfTDPAAAAAAAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAA
AAAAABk2eP1dxNpwLzsPc17oZ8EU3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAVW+YnOxyoXY3xdYvRtHNhxMg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBF
N8N/JUMuHO5/FOehwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYc
TIPsbxdLIxxVfTDPAAAAAAAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAABk
2eP1dxNpwLzsPc17oZ8EU3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAV
W+YnOxyoXY3xdYvRtHNhxMg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JUM
uHO5/FOehwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbxd
LIxxVfTDPAAAAAAAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAABk2eP1dxN
pwLzsPc17oZ8EU3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVW+YnOxy
oXY3xdYvRtHNhxMg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JUMuHO5/FO
ehwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbxdLIxxVfT
DPAAAAAAAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAABk2eP1dxNpwLzsPc
17oZ8EU3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVW+YnOxyoXY3xdY
vRtHNhxMg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JUMuHO5/FOehwAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbxdLIxxVfTDPAAAAA
AAAAAGTZ4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAABVb5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAABk2eP1dxNpwLzsPc17oZ8EU
3w38lQy4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVW+YnOxyoXY3xdYvRtHNhx
Mg+xvF0sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JUMuHO5/FOehwAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbxdLIxxVfTDPAAAAAAAAAAGT
Z4/V3E2nAvOw9zXuhnwRTfDfyVDLhzufxTnocAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABV
b5ic7HKhdjfF1i9G0c2HEyD7G8XSyMcVX0wzwAAAAAAAAABk2eP1dxNpwLzsPc17oZ8EU3w38lQy
4c7n8U56HAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVW+YnOxyoXY3xdYvRtHNhxMg+xvF0
sjHFV9MM8AAAAAAAAAAZNnj9XcTacC87D3Ne6GfBFN8N/JUMuHO5/FOehwAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAFVvmJzscqF2N8XWL0bRzYcTIPsbxdLIxxVfTDPAAAAAAAAAAGbJu7rXiv
rL3fu691IPaY9eCL18Eq4dCrJwf7ra3bwRqw1nHt2bP+NnDwcHDxcXFxbduzZw8OzbxbduzZs27X
43FOapTjJL/zGH8h/wCP/cf8h/2kfSbylZXdOvVjGEkkelN/IRmj/Jf9j/JZYRmmjySywjNGP+Qh
GKN2kOZHKmJ9ax99B+wXfJLrKec0v9X4f69fdLvYGkOZHKmJ9ax989gu+SXWU84/V+H+vX3S72Bp
DmRypifWsffPYLvkl1lPOP1fh/r190u9gaQ5kcqYn1rH3z2C75JdZTzj9X4f69fdLvYGkOZHKmJ9
ax989gu+SXWU84/V+H+vX3S72BpDmRypifWsffPYLvkl1lPOP1fh/r190u9gaQ5kcqYn1rH3z2C7
5JdZTzj9X4f69fdLvYGkOZHKmJ9ax989gu+SXWU84/V+H+vX3S72BpDmRypifWsffPYLvkl1lPOP
1fh/r190u9gaQ5kcqYn1rH3z2C75JdZTzj9X4f69fdLvYGkOZHKmJ9ax989gu+SXWU84/V+H+vX3
S72BpDmRypifWsffPYLvkl1lPOP1fh/r190u9gaQ5kcqYn1rH3z2C75JdZTzj9X4f69fdLvYGkOZ
HKmJ9ax989gu+SXWU84/V+H+vX3S72BpDmRypifWsffPYLvkl1lPOP1fh/r190u9gaQ5kcqYn1rH
3z2C75JdZTzj9X4f69fdLvYGkOZHKmJ9ax989gu+SXWU84/V+H+vX3S72BpDmRypifWsffPYLvkl
1lPOP1fh/r190u9gaQ5kcqYn1rH3z2C75JdZTzj9X4f69fdLvYGkOZHKmJ9ax989gu+SXWU84/V+
H+vX3S72BpDmRypifWsffPYLvkl1lPOP1fh/r190u9gaQ5kcqYn1rH3z2C75JdZTzj9X4f69fdLv
YGkOZHKmJ9ax989gu+SXWU84/V+H+vX3S72BpDmRypifWsffPYLvkl1lPOP1fh/r190u9gaQ5kcq
Yn1rH3z2C75JdZTzj9X4f69fdLvYGkOZHKmJ9ax989gu+SXWU84/V+H+vX3S72BpDmRypifWsffP
YLvkl1lPOP1fh/r190u9gaQ5kcqYn1rH3z2C75JdZTzj9X4f69fdLvYGkOZHKmJ9ax989gu+SXWU
84/V+H+vX3S72BpDmRypifWsffPYLvkl1lPOP1fh/r190u9gaQ5kcqYn1rH3z2C75JdZTzj9X4f6
9fdLvYGkOZHKmJ9ax989gu+SXWU84/V+H+vX3S72BpDmRypifWsffPYLvkl1lPOP1fh/r190u9ga
Q5kcqYn1rH3z2C75JdZTzj9X4f69fdLvYGkOZHKmJ9ax989gu+SXWU84/V+H+vX3S72BpDmRypif
WsffPYLvkl1lPOP1fh/r190u9gaQ5kcqYn1rH3z2C75JdZTzj9X4f69fdLvYGkOZHKmJ9ax989gu
+SXWU84/V+H+vX3S72BpDmRypifWsffPYLvkl1lPOP1fh/r190u9goRIJhPiNhXbcX7PiDdK3XX4
4zUXcrIVx2rZwcVXaNlTxRXZW7Kutq+Lj4NvFwf74f8AXD/f7s2cXDt27P5xbP7KWFnVtZYy1IQ/
sY/3/Iwm/wDcsYwUvijiKy9cqyVrSaaMssIyx6dOpRj/AH+/3/8ANaSnNGH8jD/7Qh0f+Yf3+wj/
ACjjuVcAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAB/9k=`

var mimetype = `application/epub+zip`

var ncx = `<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE ncx
  PUBLIC '-//NISO//DTD ncx 2005-1//EN'
  'http://www.daisy.org/z3986/2005/ncx-2005-1.dtd'>
<ncx version="2005-1" xml:lang="en-US" xmlns="http://www.daisy.org/z3986/2005/ncx/">
	<docTitle>
		<text>{{.Author}}</text>
	</docTitle>
	<docAuthor>
		<text>{{.Title}}</text>
	</docAuthor>
	<navMap>
		<navPoint class="" id="navpoint-1" playOrder="0">
			<navLabel>
				<text>Cover</text>
			</navLabel>
			<content src="cover.xhtml"/>
		</navPoint>
        {{range $i, $c := .Chapters}}
		<navPoint class="" id="navpoint-{{add $i 2}}" playOrder="0">
			<navLabel>
				<text>{{$c.Title}}</text>
			</navLabel>
			<content src="{{$c.Filename}}"/>
		</navPoint>
        {{end}}
	</navMap>
</ncx>`

var stylesheet = `html, body {
    margin: 0;
    padding: 0;
}
.epub {
    font-size: 1em;
    line-height: 1.25em;
    text-align: left;
}
aside {
    margin: 1em;
    padding: 0;
}
p {
    font-size: 1em;
    margin: 1em 0;
    padding: 0;
}
h1, h2, h3, h4, h5, h6 {
    font-size: 1em;
    margin: 1em 0;
    padding: 0;
    text-align: center;
}
hr {
    display: none;
}
h1 + p:first-letter,
p.break + p:first-letter {
    font-size: 1.5em;
    font-weight: bold;
}
a {
    color: #00f;
}
blockquote {
    margin: 1em;
    padding: 0;
}
.footnote {
    font-style: italic;
}
.footnote em {
    font-style: normal;
}`

var tableOfContents = `<?xml version="1.0" encoding="UTF-8"?>
<html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops">
<head>
<title>{{.Title}}</title>
{{range .Styles}}
<link rel="stylesheet" href="{{.Filename}}" type="text/css"/>
{{end}}
<meta charset="utf-8"/>
</head>
<body class="toc">
<section class="frontmatter TableOfContents" epub:type="frontmatter toc">
    <header>
        <h1>Contents</h1>
    </header>
    <nav xmlns:epub="http://www.idpf.org/2007/ops" epub:type="toc" id="toc">
        <ol>
            <li class="toc-BookTitlePage-rw" id="toc-titlepage">
                <a href="titlepage.xhtml">Title Page</a>
            </li>
            {{range .Chapters}}
            <li class="toc-Preface-rw" id="{{.Id}}">
                <a href="{{.Filename}}">{{.Title}}</a>
            </li>
            {{end}}
        </ol>
    </nav>
</section>
</body>
</html>`

var titlePage = `<?xml version="1.0" encoding="UTF-8"?>
<html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops">
<head>
<title>{{.Title}}</title>
{{range .Styles}}
<link rel="stylesheet" href="{{.Filename}}" type="text/css"/>
{{end}}
<meta charset="utf-8"/>
</head>
<body style="text-align: center;">
<div class="BookTitlePage-rw" epub:type="frontmatter titlepage">
<img src="{{.Cover}}" alt="Title Page" title="Title Page"/>
</div>
</body>
</html>`

