package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {

	//ci.Addr = r.Host
	//ci.Token = ""
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("WTF dude, error parsing your template.")
	}

	// Render the template, writing to `w`.
	t.Execute(w, r.Host)
}

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {

	type ConnInfo struct {
		Addr, Token string
	}

	var ci ConnInfo

	ci.Addr = r.Host

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		ci.Token = fmt.Sprintf("%x", h.Sum(nil))

		t, err := template.ParseFiles("templates/upload.html")
		if err != nil {
			log.Fatal("WTF dude, error parsing your template.")

		}
		t.Execute(w, ci)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
		//fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./subs/upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		//TODO convert sub
		//execCmd("iconv", "-f", "ISO-8859-2", "-t", "UTF-8", "subs/upload/" + handler.Filename, ">", "subs/ro/" + handler.Filename)
		execCmd("sh", "-c", "iconv -f ISO-8859-2 -t UTF-8 subs/upload/"+handler.Filename+" > subs/ro/"+
			handler.Filename)
		http.Redirect(w, r, "/subs/ro/", http.StatusFound)
		//execCmd("sleep", "5")
		//execCmd("ping","google.com")
		//cmd := exec.Command("sleep", "5")
		//err = cmd.Start()
		//if err != nil {
		//	//log.Fatal(err)
		//	log.Println(err)
		//}
		//log.Printf("Waiting for command to finish...")
		//err = cmd.Wait()
		//log.Printf("Command finished with error: %v", err)

	}
}

//func execCmd2(name string, arg ...string) {
//	log.Println("Executing:", name)
//	cmd := exec.Command(name, arg...)
//	err := cmd.Start()
//	if err != nil {
//		//log.Fatal(err)
//		log.Println(err)
//	}
//	log.Printf("Waiting for command to finish...")
//	err = cmd.Wait()
//	log.Printf("Command finished with error: %v", err)
//}

func execCmd(name string, arg ...string) {
	//var cmdOut []byte
	//var err error
	//cmd := "python"
	//args := []string{"/K","echo","relay.py", string(rs)}
	//args := []string{"relay.py", strconv.Itoa(arg)}
	log.Println("Executing:", name)
	cmdOut, err := exec.Command(name, arg...).Output()
	if err != nil {
		log.Println(err)
		//os.Exit(1)
	}
	log.Println("execCmd response:", string(cmdOut))
}

func main() {

	http.Handle("/subs/ro/", http.StripPrefix("/subs/ro/", http.FileServer(http.Dir("subs/ro"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}
