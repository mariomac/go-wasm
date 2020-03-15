package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/mariomac/gorrazo/pkg/draw"
)

func main() {
	c := draw.GetCanvas("theCanvas",
		draw.FullScreen(true))

	c.ClearRect(0, 0, c.GetWidth(), c.GetHeight())
	txtPos := float64(c.GetHeight())
	counter := float64(0)
	img := draw.HTMLImageElement("./img/virus.png")
	for {
		for col := 0; col < 256; col++ {
			c.FillStyle(fmt.Sprintf("#%x%x%x", 255-col, 64+col/2, col))
			c.FillRect(0, c.GetHeight()*float64(col)/255, c.GetWidth(), c.GetHeight()/256)
		}
		c.FillStyle("yellow")
		c.StrokeStyle("black")
		c.Font("40px Arial")
		for l, line := range txt {
			c.FillText(line, 0, txtPos+float64(l)*50)
			c.StrokeText(line, 0, txtPos+float64(l)*50)
		}
		c.DrawImageD(img, 20, 20, 200, 200)
		c.DrawImage(img, 100, 100)
		c.Save()
		c.Translate(360, 160)
		c.Rotate(counter/10)
		c.Translate(-360, -160)
		c.DrawImageD(img, 360, 160, 120, 120)
		c.Restore()
		txtPos -= 1
		counter++
		time.Sleep(20 * time.Millisecond)
	}

}

var txt = strings.Split(`Yo ya fui yo
Yo ya fui yo
Yo ya fui
A Cangas del morrazo
Y menos mal que no llovía
Vámonos vámonos
Vámonos nena
Al otro lado de la ría
Viera gente de Pontevedra
Y viera gente de marín
Viera a suso de moaña
Que pasara por allí
Ya comiera y ya bebiera
Y menos mal que no lloviera.
Yo ya fui yo
Yo ya fui yo
Yo ya fui
A Cangas del morrazo
Y menos mal que no llovía
Vámonos vámonos
Vámonos nena
Al otro lado de la ría
Viera mujeres con carritos
Viera gente marinera
Viera barcos y gaviotas
Y menos mal que no lloviera
Viera chocos y fanecas
Y gamelas y bateas.
Viera vigo desde lejos
Y menos mal que no lloviera.
Yo ya fui yo
Yo ya fui yo
Yo ya fui
A Cangas del morrazo
Y menos mal que no llovía
Vámonos vámonos
Vámonos nena
Al otro lado de la ría
Vámonos vámonos
Vámonos nena
Al otro lado de la ría.`, "\n")
