package main

import (
	"github.com/electricbubble/guia2"
	"log"
	"time"
)

func main() {
	// driver, err := guia2.NewDriver(guia2.NewEmptyCapabilities(), "http://localhost:6790/wd/hub")
	driver, err := guia2.NewDriver(nil, "http://localhost:6790/wd/hub")
	checkErr(err)

	// fmt.Println(driver.Source())
	// return

	deviceSize, err := driver.DeviceSize()
	checkErr(err)

	var startX, startY, endX, endY int
	startX = deviceSize.Width / 2
	startY = deviceSize.Height / 2
	endX = startX
	endY = startY / 2
	err = driver.Swipe(startX, startY, endX, endY)
	checkErr(err)

	var startPoint, endPoint guia2.PointF
	startPoint = guia2.PointF{X: float64(startX), Y: float64(startY)}
	endPoint = guia2.PointF{X: startPoint.X, Y: startPoint.Y * 1.6}
	err = driver.SwipePointF(startPoint, endPoint)
	checkErr(err)

	element, err := driver.FindElement(guia2.BySelector{UiAutomator: "new UiSelector().className(\"android.view.ViewGroup\");"})
	checkErr(err)

	elem, err := element.FindElement(guia2.BySelector{UiAutomator: "new UiSelector().className(\"android.widget.LinearLayout\").index(6);"})
	checkErr(err)

	rect, err := elem.Rect()
	checkErr(err)

	x := rect.X + int(float64(rect.Width)*2)
	y := rect.Y + rect.Height/2
	err = driver.Tap(x, y)
	checkErr(err)

	time.Sleep(time.Millisecond * 600)

	element, err = driver.FindElement(guia2.BySelector{UiAutomator: "new UiSelector().text(\"科技\");"})
	checkErr(err)

	time.Sleep(time.Millisecond * 1200)

	err = element.Click()
	checkErr(err)

	time.Sleep(time.Second * 1)

	err = driver.Swipe(startX, startY, endX, endY/2)
	checkErr(err)

	elements, err := driver.FindElements(guia2.BySelector{ResourceIdID: "cn.xuexi.android:id/general_card_image_id"})
	checkErr(err)
	time.Sleep(time.Second * 1)

	elem = elements[len(elements)-1]
	checkErr(elem.Click())
	time.Sleep(time.Second * 1)

	err = driver.ScrollTo(guia2.BySelector{ContentDescription: "点赞"})
	checkErr(err)

	element, err = driver.FindElement(guia2.BySelector{ClassName: "android.widget.TextView"})
	checkErr(err)
	checkErr(element.Click())
	time.Sleep(time.Millisecond * 500)

	elem, err = driver.FindElement(guia2.BySelector{ClassName: "android.widget.EditText"})
	checkErr(err)

	err = elem.SendKeys("test " + time.Now().Format("2006-01-02 15:04:05"))
	checkErr(err)
	time.Sleep(time.Second * 1)

	element, err = driver.FindElement(guia2.BySelector{UiAutomator: "new UiSelector().text(\"取消\");"})
	checkErr(err)
	checkErr(element.Click())

	// element, err = driver.FindElement(guia2.BySelector{ResourceIdID: "cn.xuexi.android:id/TOP_LAYER_VIEW_ID"})
	// checkErr(err)
	// elemBack, err := element.FindElement(guia2.BySelector{ClassName: "android.widget.ImageView"})
	// checkErr(err)

	// screenshot, err := elem.Screenshot()
	// checkErr(err)
	// ioutil.WriteFile("/path/Desktop/e1.png", screenshot.Bytes(), 0600)

}

func checkErr(err error, msg ...string) {
	if err == nil {
		return
	}

	var output string
	if len(msg) != 0 {
		output = msg[0] + " "
	}
	output += err.Error()
	log.Fatalln(output)
}
