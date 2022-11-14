package main

import (
	"daily/demo-project/onenote/service"
	"encoding/xml"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/grokify/html-strip-tags-go"
	"io/ioutil"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var (
	xmlContentMap = new(sync.Map) //key:datetime val:multi task
	taskMap       = new(sync.Map) //key:datetime val:multi task
	globalApp     = app.NewWithID("reminder-test")
)

func init() {
	//err := os.Setenv("FYNE_FONT", fmt.Sprintf("%s", "resource/ttf/SourceHanSerifCN-Regular.ttf"))
	//if nil != err {
	//	panic(err)
	//}
	//fmt.Println(os.Getenv("FYNE_FONT"))
}

func main() {
	var start service.StartChan
	start.Start = make(chan bool)

	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()

		for {
			exeName := "D:\\go-workspace\\src\\onenote\\VanillaAddIn\\VanillaConsole\\bin\\Release\\VanillaConsole.exe"
			invokeNoteExe(exeName)
			fmt.Println(fmt.Sprintf("note is write at %s", time.Now().String()))

			fileName := "D:/content.xml"
			readXmlContent(fileName)
			fmt.Println(fmt.Sprintf("note is read at %s", time.Now().String()))

			if !start.IsClosed {
				start.Start <- true
				close(start.Start)
			}

			<-ticker.C
		}
	}()

	go func() {
		validateDatetime()
	}()

	if <-start.Start {
		winFromShow(true, "ALL")
		start.IsClosed = true
	}

	globalApp.Run()
}

// invoke exe read note content
func invokeNoteExe(exeName string) {
	cmd := exec.Command(exeName)
	if err := cmd.Run(); nil != err {
		fmt.Println(`exec error:`, err.Error())
		//panic(err.Error())
	}
}

// read xml content
func readXmlContent(fileName string) {
	// reset map
	// fix bug list 6
	tStatus := make(map[string]bool)
	xmlContentMap.Range(func(key, value interface{}) bool {
		tStatus[key.(string)] = value.(service.TaskInfo).Status
		xmlContentMap.Delete(key)
		return true
	})

	data, err := ioutil.ReadFile(fileName)
	if nil != err {
		fmt.Println(`io read file error:`, err.Error())
		return
	}

	var oneContent service.Page
	if err = xml.Unmarshal(data, &oneContent); nil != err {
		fmt.Println(`xml unmarshal error:`, err.Error())
		return
	}

	//fmt.Println(fmt.Sprintf("%+v", oneContent))
	for _, outLine := range oneContent.Outline {
		for _, oe := range outLine.OEChildren.OE {

			//fmt.Println("easily reminder", strip.StripTags(oe.T.Text))
			// table reminder : recommend
			for _, row := range oe.Table.Row {
				if 2 <= len(row.Cell.OEChildren.OE) {
					//fmt.Println(fmt.Sprintf("task is:%s, time is:%s", row.Cell.OEChildren.OE[0].T, row.Cell.OEChildren.OE[1].T))

					t := strip.StripTags(row.Cell.OEChildren.OE[1].T)
					if !strings.Contains(t, "!") {
						continue
					}

					var tmp service.TaskInfo

					tmp.Status = true
					if v, ok := tStatus[t]; ok {
						tmp.Status = v
					}
					tmp.Name = append(tmp.Name, strip.StripTags(row.Cell.OEChildren.OE[0].T))
					xmlContentMap.Store(t, tmp)
				}
			}
		}
	}
	taskMap.Range(func(key, value interface{}) bool {
		taskMap.Delete(key)
		return true
	})
	xmlContentMap.Range(func(key, value interface{}) bool {
		taskMap.Store(key, value)
		fmt.Println(key, value)
		return true
	})
}

// show windows form
func winFromShow(isMain bool, key string) {
	w := globalApp.NewWindow("Reminder from onenote")

	var data [][]string

	taskMap.Range(func(key, value interface{}) bool {

		for _, v := range value.(service.TaskInfo).Name {
			data = append(data, []string{v, key.(string)})
		}

		return true
	})

	fmt.Println(fmt.Sprintf("%+v", data))

	table := widget.NewTable(func() (int, int) {
		return len(data), len(data[0])
	}, func() fyne.CanvasObject {
		label := widget.NewLabel("task content")
		return label
	}, func(id widget.TableCellID, object fyne.CanvasObject) {
		if key == data[id.Row][1] {
			object.(*widget.Label).TextStyle = fyne.TextStyle{Bold: true}
		}
		object.(*widget.Label).SetText(data[id.Row][id.Col])
	})

	// set col width
	table.SetColumnWidth(0, 300)

	w.SetContent(table)
	w.Resize(fyne.NewSize(570, 350))

	if isMain {
		w.SetCloseIntercept(func() {
			w.Hide()
		})
	}
	w.Show()
	// todo need think whether status false
}

func validateDatetime() {
	tick := time.Tick(20 * time.Second)
	for {
		<-tick
		t := time.Now()
		ymdhms := fmt.Sprintf("!%s", t.Format("2006-01-02 15:04:05"))
		ymdhm := fmt.Sprintf("!%s", t.Format("2006-01-02 15:04"))
		ymd := fmt.Sprintf("!%s", t.Format("2006-01-02"))
		//hms := fmt.Sprintf("!%s", t.Format("15:04:05")) // interval 20 seconds, maybe can not check hms
		hm := fmt.Sprintf("!%s", t.Format("15:04"))
		//fmt.Println(time.Now(), fmt.Sprintf("%+v", taskMap))

		if v, ok := taskMap.Load(ymdhms); ok {
			tInfo := v.(service.TaskInfo)
			if tInfo.Status {
				winFromShow(false, ymdhms)
				tInfo.Status = false
				taskMap.Store(ymdhms, tInfo)
			}
		}

		if v, ok := taskMap.Load(ymdhm); ok {
			tInfo := v.(service.TaskInfo)
			if tInfo.Status {
				winFromShow(false, ymdhm)
				tInfo.Status = false
				taskMap.Store(ymdhm, tInfo)
			}
		}

		if v, ok := taskMap.Load(ymd); ok {
			tInfo := v.(service.TaskInfo)
			if tInfo.Status {
				winFromShow(false, ymd)
				tInfo.Status = false
				taskMap.Store(ymd, tInfo)
			}
		}

		// handle hm
		taskMap.Range(func(key, value interface{}) bool {
			if !strings.HasPrefix(key.(string), hm) {
				return true // need continue iteration
			}
			tInfo := value.(service.TaskInfo)
			if tInfo.Status {
				winFromShow(false, key.(string))
				tInfo.Status = false
				taskMap.Store(key.(string), tInfo)
			}
			return true
		})
	}
}
