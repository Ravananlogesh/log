package helpers

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"api-log/genpkg"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

const (

	Elog      = 0
	Statement = 1
	Details   = 2
)

type HelperStruct struct {
	Sid       string
	Reference string
}

func (h *HelperStruct) Init() {
	KeyValue := ""
	session := uuid.NewV4()
	sessionSHA256 := session.String()
	KeyValue = strings.ReplaceAll(sessionSHA256, "-", "")
	h.Sid = KeyValue
}

func (h *HelperStruct) SetReference(pReference interface{}) {
	h.Reference = fmt.Sprintf("%v", pReference)
}

func (h *HelperStruct) RemoveReference() {
	h.Reference = ""
}

func ErrReturn(pErr error) error {
	lErr := ""
	lPc, lFile, lLine, _ := runtime.Caller(1)
	lFuncname := runtime.FuncForPC(lPc).Name()
	lStrArray := strings.Split(lFile, "/")
	lFilename := lStrArray[len(lStrArray)-2] + "/" + lStrArray[len(lStrArray)-1]

	lErr = lFilename + " @@ " + lFuncname + " @@ ln " + fmt.Sprintf("%d", lLine) + " @@ "

	if strings.Contains(pErr.Error(), " @@ ") && strings.Contains(pErr.Error(), " @@ ln ") {
		lErr = ""
	} else {

		lErr = lFilename + " @@ " + lFuncname + " @@ ln " + fmt.Sprintf("%d", lLine) + " @@ "
	}
	return errors.New(lErr + pErr.Error())
}

func (h *HelperStruct) Log(pDebugLevel int, pMsg ...interface{}) {
	// find the mata data of printed value
	lPc, lFile, lLine, _ := runtime.Caller(1)
	lFuncname := runtime.FuncForPC(lPc).Name()

	//read the value from toml

	lConfigFile := genpkg.ReadTomlConfig("./toml/debug.toml")
	lLevel := fmt.Sprintf("%v", lConfigFile.(map[string]interface{})["LogCategory"])
	intlevel, err := strconv.Atoi(lLevel)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	lReference := fmt.Sprintf("%v", lConfigFile.(map[string]interface{})["LogReference"])

	//file name over writting
	lStrArray := strings.Split(lFile, "/")
	lFilename := lStrArray[len(lStrArray)-2] + "/" + lStrArray[len(lStrArray)-1]

	//find the Debug level
	if (pDebugLevel <= intlevel && intlevel != 0) || pDebugLevel == Elog {
		str := fmt.Sprintf("%v", pMsg)
		lFinal := str[1 : len(str)-1]

		//check the sid will be set
		if strings.EqualFold(h.Sid, "") {
			log.Println("Set the id before <-- debug := new(helpers.HelperStruct) after debug.Init(reference value) -->")
		} else {

			//print the O/P based on reference value
			if lReference == "" || strings.EqualFold(lReference, h.Reference) {
				if strings.Contains(lFinal, " @@ ") && strings.Contains(lFinal, "@@ ln") {
					log.Println("@@", h.Sid, "@@ (", h.Reference, ") @@", lFinal)
				} else {
					log.Println("@@", h.Sid, "@@ (", h.Reference, ") @@", lFilename, "@@", lFuncname, "@@ ln", lLine, "@@", lFinal)
				}
			}
		}

	}
}

func ErrPrint(err error) string {
	lStrArray := strings.Split(err.Error(), "@@")
	return lStrArray[len(lStrArray)-1]
}
